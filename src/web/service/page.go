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
}

type PageData struct {
	Title          string
	Maps           []MapOption
	SelectedMap    string
	ExpandedMap    string
	MethodGroups   []MethodGroup
	SelectedMethod string
	ExpandedMethod string
	SelectedRun    string
	Points         points.Points2D
	PointsJSON     string
}

type SelectionState struct {
	SelectedMap    string
	ExpandedMap    string
	SelectedMethod string
	ExpandedMethod string
	SelectedRun    string
}

type PageService struct {
	repo  repository.Reader
	title string
}

func NewPageService(repo repository.Reader) *PageService {
	return &PageService{
		repo:  repo,
		title: defaultTitle,
	}
}

func (s *PageService) BuildPage(ctx context.Context, selection SelectionState) (PageData, error) {
	selectedMap := strings.TrimSpace(selection.SelectedMap)
	expandedMap := strings.TrimSpace(selection.ExpandedMap)
	selectedMethod := strings.TrimSpace(selection.SelectedMethod)
	expandedMethod := strings.TrimSpace(selection.ExpandedMethod)
	selectedRun := strings.TrimSpace(selection.SelectedRun)

	availableMapIDs, err := s.repo.ListMaps(ctx)
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

	loadedPoints, err := s.repo.LoadPoints(ctx, selectedMap)
	if err != nil {
		return PageData{}, fmt.Errorf("load points for map %q: %w", selectedMap, err)
	}

	pointsJSON, err := json.Marshal(loadedPoints)
	if err != nil {
		return PageData{}, fmt.Errorf("marshal points for map %q: %w", selectedMap, err)
	}

	loadedRuns, err := s.repo.ListRuns(ctx, selectedMap)
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

	for _, run := range runs {
		if run.RunID == selectedRun && run.Method == selectedMethod {
			data.SelectedRun = selectedRun
			break
		}
	}

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
