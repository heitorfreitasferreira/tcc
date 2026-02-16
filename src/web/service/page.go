package service

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"

	"tcc/points"
	"tcc/web/repository"
)

const defaultTitle = "TCC - Visualizador"

var preferredMethods = map[string]int{
	"aco":        0,
	"pso":        1,
	"ga":         2,
	"bruteforce": 3,
}

type MapOption struct {
	ID string
}

type MethodGroup struct {
	Method string
	Runs   []RunOption
}

type RunOption struct {
	RunID        string
	RunHash      string
	Method       string
	Seed         int64
	BestMakespan float64
	Iterations   int
}

type PageData struct {
	Title                   string
	Maps                    []MapOption
	SelectedMap             string
	ExpandedMap             string
	MethodGroups            []MethodGroup
	SelectedMethod          string
	ExpandedMethod          string
	SelectedRun             string
	SelectedRunIterations   int
	SelectedRunBestMakespan float64
	EvolutionFramesCount    int
	Points                  points.Points2D
	PointsJSON              string
	EvolutionJSON           string
}

type SelectionState struct {
	SelectedMap    string
	ExpandedMap    string
	SelectedMethod string
	ExpandedMethod string
	SelectedRun    string
}

type PageService struct {
	mapsRepo      repository.MapsReader
	summaryRepo   repository.SummaryReader
	evolutionRepo repository.EvolutionReader
	title         string
}

func NewPageService(mapsRepo repository.MapsReader, summaryRepo repository.SummaryReader, evolutionRepo repository.EvolutionReader) *PageService {
	return &PageService{
		mapsRepo:      mapsRepo,
		summaryRepo:   summaryRepo,
		evolutionRepo: evolutionRepo,
		title:         defaultTitle,
	}
}

type evolutionPayload struct {
	RunID      string                      `json:"run_id"`
	Iterations int                         `json:"iterations"`
	Frames     []repository.EvolutionFrame `json:"frames"`
}

func (s *PageService) BuildPage(ctx context.Context, selection SelectionState) (PageData, error) {
	selectedMap := strings.TrimSpace(selection.SelectedMap)
	expandedMap := strings.TrimSpace(selection.ExpandedMap)
	selectedMethod := strings.TrimSpace(selection.SelectedMethod)
	expandedMethod := strings.TrimSpace(selection.ExpandedMethod)
	selectedRun := strings.TrimSpace(selection.SelectedRun)

	availableMapIDs, err := s.mapsRepo.ListMaps(ctx)
	if err != nil {
		return PageData{}, fmt.Errorf("load available maps: %w", err)
	}

	data := PageData{
		Title: s.title,
		Maps:  toMapOptions(availableMapIDs),
	}

	if selectedMap == "" || !slices.Contains(availableMapIDs, selectedMap) {
		return data, nil
	}

	loadedPoints, err := s.mapsRepo.LoadPoints(ctx, selectedMap)
	if err != nil {
		return PageData{}, fmt.Errorf("load points for map %q: %w", selectedMap, err)
	}

	pointsJSON, err := json.Marshal(loadedPoints)
	if err != nil {
		return PageData{}, fmt.Errorf("marshal points for map %q: %w", selectedMap, err)
	}

	loadedRuns, err := s.summaryRepo.ListRuns(ctx, selectedMap)
	if err != nil {
		return PageData{}, fmt.Errorf("load runs for map %q: %w", selectedMap, err)
	}

	runs := toRunOptions(loadedRuns)
	methodGroups := groupRunsByMethod(runs)

	data.SelectedMap = selectedMap
	if expandedMap == selectedMap {
		data.ExpandedMap = selectedMap
	}
	data.MethodGroups = methodGroups
	data.Points = loadedPoints
	data.PointsJSON = string(pointsJSON)

	if data.ExpandedMap == "" {
		expandedMethod = ""
	}

	if selectedMethod == "" && expandedMethod == "" && selectedRun == "" {
		return data, nil
	}

	if selectedRun != "" {
		for _, run := range runs {
			if run.RunID != selectedRun {
				continue
			}

			if selectedMethod == "" {
				selectedMethod = run.Method
			}

			if expandedMethod == "" {
				expandedMethod = run.Method
			}

			break
		}
	}

	selectedMethodValid := false
	expandedMethodValid := false
	for _, group := range methodGroups {
		if group.Method == selectedMethod {
			selectedMethodValid = true
		}

		if group.Method == expandedMethod {
			expandedMethodValid = true
		}
	}

	if selectedMethodValid {
		data.SelectedMethod = selectedMethod
	} else {
		selectedMethod = ""
		selectedRun = ""
	}

	if expandedMethodValid {
		data.ExpandedMethod = expandedMethod
	}

	if selectedRun == "" || selectedMethod == "" {
		return data, nil
	}

	var selectedRunData RunOption
	selectedRunValid := false
	for _, run := range runs {
		if run.RunID == selectedRun && run.Method == selectedMethod {
			selectedRunData = run
			selectedRunValid = true
			data.SelectedRun = selectedRun
			break
		}
	}

	if !selectedRunValid {
		return data, nil
	}

	loadedEvolution, err := s.evolutionRepo.LoadEvolution(ctx, selectedRun)
	if err != nil {
		return PageData{}, fmt.Errorf("load evolution for run %q: %w", selectedRun, err)
	}

	loadedGraph, err := s.mapsRepo.LoadGraph(ctx, selectedMap)
	if err != nil {
		return PageData{}, fmt.Errorf("load graph for map %q: %w", selectedMap, err)
	}

	nodeLimit := len(loadedPoints)
	if graphSize := len(loadedGraph); graphSize > 0 && graphSize < nodeLimit {
		nodeLimit = graphSize
	}

	normalizedFrames := normalizeEvolutionFrames(loadedEvolution, nodeLimit)
	payload := evolutionPayload{
		RunID:      selectedRun,
		Iterations: selectedRunData.Iterations,
		Frames:     normalizedFrames,
	}

	encodedEvolution, err := json.Marshal(payload)
	if err != nil {
		return PageData{}, fmt.Errorf("marshal evolution for run %q: %w", selectedRun, err)
	}

	data.SelectedRunIterations = selectedRunData.Iterations
	data.SelectedRunBestMakespan = selectedRunData.BestMakespan
	data.EvolutionFramesCount = len(normalizedFrames)
	data.EvolutionJSON = string(encodedEvolution)

	return data, nil
}

func (s *PageService) ResolveMapSelection(current SelectionState, targetMap string) SelectionState {
	targetMap = strings.TrimSpace(targetMap)
	currentMap := strings.TrimSpace(current.SelectedMap)
	expandedMap := strings.TrimSpace(current.ExpandedMap)

	selectedMap := currentMap
	if selectedMap == "" {
		selectedMap = targetMap
	}

	if targetMap != "" {
		if currentMap == "" || targetMap != currentMap {
			selectedMap = targetMap
			expandedMap = targetMap
		} else if expandedMap == targetMap {
			expandedMap = ""
		} else {
			expandedMap = targetMap
		}
	}

	return SelectionState{
		SelectedMap: selectedMap,
		ExpandedMap: expandedMap,
	}
}

func (s *PageService) ResolveMethodSelection(current SelectionState, targetMethod string) SelectionState {
	selectedMethod := strings.TrimSpace(current.SelectedMethod)
	expandedMethod := strings.TrimSpace(current.ExpandedMethod)
	targetMethod = strings.TrimSpace(targetMethod)

	if targetMethod != "" {
		if expandedMethod == targetMethod {
			expandedMethod = ""
			if selectedMethod == targetMethod {
				selectedMethod = ""
			}
		} else {
			selectedMethod = targetMethod
			expandedMethod = targetMethod
		}
	}

	selectedMap := strings.TrimSpace(current.SelectedMap)

	return SelectionState{
		SelectedMap:    selectedMap,
		ExpandedMap:    selectedMap,
		SelectedMethod: selectedMethod,
		ExpandedMethod: expandedMethod,
	}
}

func (s *PageService) ResolveRunSelection(mapID, method, runID string) SelectionState {
	selectedMap := strings.TrimSpace(mapID)
	selectedMethod := strings.TrimSpace(method)

	return SelectionState{
		SelectedMap:    selectedMap,
		ExpandedMap:    selectedMap,
		SelectedMethod: selectedMethod,
		ExpandedMethod: selectedMethod,
		SelectedRun:    strings.TrimSpace(runID),
	}
}

func toMapOptions(mapIDs []string) []MapOption {
	result := make([]MapOption, 0, len(mapIDs))
	for _, mapID := range mapIDs {
		result = append(result, MapOption{ID: mapID})
	}

	return result
}

func toRunOptions(records []repository.RunRecord) []RunOption {
	runs := make([]RunOption, 0, len(records))
	for _, record := range records {
		runs = append(runs, RunOption{
			RunID:        record.RunID,
			RunHash:      extractRunHash(record.RunID),
			Method:       record.Method,
			Seed:         record.Seed,
			BestMakespan: record.BestMakespan,
			Iterations:   record.Iterations,
		})
	}

	return runs
}

func extractRunHash(runID string) string {
	idx := strings.LastIndex(runID, "__")
	if idx == -1 || idx+2 >= len(runID) {
		return runID
	}

	return runID[idx+2:]
}

func normalizeEvolutionFrames(frames []repository.EvolutionFrame, nodeLimit int) []repository.EvolutionFrame {
	if len(frames) == 0 {
		return []repository.EvolutionFrame{}
	}

	normalized := make([]repository.EvolutionFrame, 0, len(frames))
	for _, frame := range frames {
		normalized = append(normalized, repository.EvolutionFrame{
			Iter:         nonNegative(frame.Iter),
			EvalCount:    nonNegative(frame.EvalCount),
			BestMakespan: frame.BestMakespan,
			BestSequence: normalizeSequence(frame.BestSequence, nodeLimit),
		})
	}

	return normalized
}

func normalizeSequence(sequence []int, nodeLimit int) []int {
	if len(sequence) == 0 || nodeLimit <= 0 {
		return []int{}
	}

	normalized := make([]int, 0, len(sequence))
	for _, index := range sequence {
		if index < 0 || index >= nodeLimit {
			continue
		}

		normalized = append(normalized, index)
	}

	return normalized
}

func nonNegative(value int) int {
	if value < 0 {
		return 0
	}

	return value
}

func groupRunsByMethod(runs []RunOption) []MethodGroup {
	grouped := make(map[string][]RunOption)

	for _, run := range runs {
		grouped[run.Method] = append(grouped[run.Method], run)
	}

	groups := make([]MethodGroup, 0, len(grouped))
	for method, methodRuns := range grouped {
		sort.Slice(methodRuns, func(i, j int) bool {
			if methodRuns[i].Seed != methodRuns[j].Seed {
				return methodRuns[i].Seed < methodRuns[j].Seed
			}

			return methodRuns[i].RunID < methodRuns[j].RunID
		})

		groups = append(groups, MethodGroup{
			Method: method,
			Runs:   methodRuns,
		})
	}

	sort.Slice(groups, func(i, j int) bool {
		rankI, okI := preferredMethods[groups[i].Method]
		rankJ, okJ := preferredMethods[groups[j].Method]

		if okI && okJ {
			return rankI < rankJ
		}

		if okI {
			return true
		}

		if okJ {
			return false
		}

		return groups[i].Method < groups[j].Method
	})

	return groups
}
