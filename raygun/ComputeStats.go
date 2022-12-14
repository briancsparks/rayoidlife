package raygun

import "sync"

// -------------------------------------------------------------------------------------------------------------------

type ComputeStatsData struct {
  Cmps        int
  Sqrts       int

  Points      int
  PointsProc  int
}

func (std *ComputeStatsData) accumulate(that *ComputeStatsData) {
  std.Cmps        += that.Cmps
  std.Sqrts       += that.Sqrts
  std.Points      += that.Points
  std.PointsProc  += that.PointsProc
}

// -------------------------------------------------------------------------------------------------------------------

type ComputeStats struct {
  data      ComputeStatsData

  AddStats *chan ComputeStatsData

  Stats    *chan chan ComputeStatsData
}

func (st *ComputeStats) addStats(std ComputeStatsData) {
  *st.AddStats <- std
}

func (st *ComputeStats) Reset() {
  st.data.Cmps        = 0
  st.data.Sqrts       = 0
  st.data.Points      = 0
  st.data.PointsProc  = 0
}

// -------------------------------------------------------------------------------------------------------------------

func StartComputeStatsAgent() *ComputeStats {

  stats   := make(chan chan ComputeStatsData)
  addData := make(chan ComputeStatsData)

  st := &ComputeStats{
    data: ComputeStatsData{
      Cmps:       0,
      Sqrts:      0,
      Points:     0,
      PointsProc: 0,
    },
    Stats:   &stats,
    AddStats: &addData,
  }
  st.start()
  return st
}

// -------------------------------------------------------------------------------------------------------------------

func (st *ComputeStats) start() {
  wg := sync.WaitGroup{}

  wg.Add(1)
  go func() {
    wg.Done()

    for {
      select {
      case d := <- *st.AddStats:
        st.data.accumulate(&d)

      case ch := <- *st.Stats:
        ch <- ComputeStatsData{
          Cmps:       st.data.Cmps,
          Sqrts:      st.data.Sqrts,
          Points:     st.data.Points,
          PointsProc: st.data.PointsProc,
        }
      }
    }
  }()
  wg.Wait()

}

// -------------------------------------------------------------------------------------------------------------------

func (st *ComputeStats) GetData() ComputeStatsData {
  result := ComputeStatsData{}
  ch := make(chan ComputeStatsData)

  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    defer wg.Done()
    result = <- ch
  }()
  *st.Stats <- ch

  wg.Wait()
  return result
}
