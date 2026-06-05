## Bock et al. (2025) — A Survey on the TSP and Its Variants in a Warehousing Context

**Reference:** Bock, S., Bomsdorf, S., Boysen, N., & Schneider, M. (2025). A survey on the Traveling Salesman Problem and its variants in a warehousing context. *European Journal of Operational Research*, 322, 1–14. doi:10.1016/j.ejor.2024.04.014

### 1. Problem and motivation

E-commerce and fast-delivery expectations have renewed interest in efficient picker routing within warehouses and distribution centers. While classical single-picker routing maps directly to the TSP and is solvable in polynomial time under the parallel-aisle (block) structure of warehouses, the diversification of warehousing processes — driven by automation, scattered storage, AMR-assisted picking, and combined stowing/picking — requires solving more complex TSP variants (clustered, generalized, prize-collecting, orienteering, etc.). The paper systematically surveys which TSP variants have meaningful warehouse applications and analyzes their computational complexity under the block-structured distance matrix.

### 2. Core method or approach

- **Comprehensive variant identification:** Started with 54 TSP variants from the routing literature, filtered through a multi-stage expert selection process (brainstorming with 7 researchers + warehousing consultant, validation by 5 independent experts) to identify 10 TSP variants with plausible warehousing use cases.
- **Structured treatment per variant:** For each selected variant, the paper provides: (i) a formal problem definition with ILP formulations and state-of-the-art solution methods, (ii) one or more warehouse use cases, (iii) computational complexity analysis in single-block (1B), two-block (2B), and multi-block (MB) layouts, and (iv) future research needs.
- **Complexity analysis framework:** Distinguishes the classical TSP on general graphs (NP-hard) from the TSP on block-structured distance matrices (often polynomial). Denotes warehouse-specific versions as W-TSP, 1B-TSP, 2B-TSP, MB-TSP.
- **New complexity proofs:** Provides three original complexity results — strong NP-hardness of 1B-TSPTW (reduction from Line-TSPTW), binary NP-hardness of 1B-OP, and binary NP-hardness of 1B-PCTSP — along with pseudo-polynomial exact algorithms for OP and PCTSP when the number of cross aisles *h* ∈ O(log n).
- **Polynomial algorithms via dynamic programming:** Surveys the DP lineage from Ratliff & Rosenthal (1983) for 1B-TSP through Roodbergen & de Koster (2001b) for 2B-TSP to Cambazard & Catusse (2018) and Pansart et al. (2018) for MB-TSP, and extends DP nesting techniques to CTSP-GCS and TSPB.

### 3. Main results

- **10 warehouse-relevant TSP variants identified** (Table 1):
  - **Polynomial:** 1B/2B/MB-TSP (Ratliff & Rosenthal, 1983; Roodbergen & de Koster, 2001b; Pansart et al., 2018), MB-CTSP-GCS (Löffler et al., 2021), MB-TSPB (Žulj et al., 2018).
  - **NP-hard:** 1B-CTSP-OCS (strongly NP-hard), 1B-GTSP (strongly NP-hard), 1B-PCTSP (binary NP-hard), 1B-OP (binary NP-hard), 1B-TSPTW (strongly NP-hard).
  - **Open complexity:** 1B/2B/MB-TSPPC, 1B/2B/MB-TRP, 1B/2B/MB-k-best TSP, 1B/2B/MB-CSP.
- **Warehouse use cases mapped:** GTSP → scattered storage picking; TSPB → combined stowing and picking; PCTSP → stowing in scattered storage; OP → picking under cutoff time constraints; CSP → RFID-based automated stock-taking; TSPTW → cold-chain preservation and team coordination via time-window decomposition; CTSP → AMR-assisted picking and multi-depot tours.
- **Performance baselines surveyed:** For each TSP variant, the paper reports representative solver performance (e.g., Concorde solves 500-node STSP in <1 min; Fischetti et al. solve GTSP with 442 nodes/89 clusters in 16.3 h heuristically in seconds).
- **Future research directions:** Open complexity statuses for TSPPC, TRP, k-best TSP, and CSP; dynamic/stochastic extensions (uncertain stowing capacity, returned products); team coordination via decomposition with k-best TSP subproblems; holistic routing-packing models for precedence constraints; family TSP with piece-level demand bookkeeping for scattered storage.

### 4. Strengths and limitations

**Strengths:**
- First survey to systematically bridge the TSP variant literature with warehouse-specific complexity analysis, combining perspectives from both routing and warehousing research communities.
- Provides three original complexity proofs and identifies four open cases, directly guiding future theoretical work.
- Each variant's warehouse use case is described concretely with references to specific warehousing processes (AMR-assisted picking, scattered storage, cold-chain preservation, etc.).
- Comprehensive scope: 54 variants screened, of which 10 are treated in depth with uniform structure.

**Limitations:**
- Restricted to single-picker, static, deterministic problems; excludes multi-picker coordination, dynamic order arrivals, and stochastic travel or stowing outcomes.
- Only the parallel-aisle block layout is considered; alternative warehouse layouts (flying-V, fishbone, discrete cross-aisle) are explicitly excluded.
- The variant selection process, despite the multi-expert validation attempt, is acknowledged by the authors as biased by subjective assessment.
- Team coordination is only sketched via decomposition approaches (k-best TSP, time-window decomposition) without empirical evaluation.

### 5. Practical takeaway for researchers

The block structure of warehouses is a powerful structural property: several TSP variants that are NP-hard on general graphs become polynomial (e.g., MB-TSP, MB-CTSP-GCS, MB-TSPB), enabling exact optimization for realistically sized instances. When designing algorithms for warehouse routing problems, researchers should first check whether the block structure yields tractability before resorting to heuristics. For the NP-hard variants (GTSP, PCTSP, OP, TSPTW, CTSP-OCS), the paper identifies exact pseudo-polynomial approaches for some (OP, PCTSP when *h* is bounded) and highlights open complexity cases (TSPPC, TRP, CSP) as direct targets for future complexity-theoretic work. Decomposition methods that use efficient single-worker W-TSP variants as subproblems are a promising direction for multi-picker coordination without sacrificing picking performance.
