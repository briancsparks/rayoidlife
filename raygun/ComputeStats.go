package raygun

import "sync"

// -------------------------------------------------------------------------------------------------------------------

type ComputeStats struct {
  Cmps      int
  Sqrts     int

  AddCmp   *chan int
  AddSqrt  *chan int

  Stats    *chan chan ComputeStats
}

func (st *ComputeStats) addCmp(n int) {
  *st.AddCmp <- n
}

func (st *ComputeStats) addSqrt(n int) {
  *st.AddSqrt <- n
}

func (st *ComputeStats) Reset() {
  st.Cmps = 0
  st.Sqrts = 0
}

// -------------------------------------------------------------------------------------------------------------------

func StartComputeStatsAgent() *ComputeStats {
  addCmp  := make(chan int, 1000)
  addSqrt := make(chan int, 1000)

  stats   := make(chan chan ComputeStats)

  st := &ComputeStats{
    AddSqrt: &addSqrt,
    AddCmp:  &addCmp,
    Stats:   &stats,
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

    var n int

    for {
      select {
      case n = <- *st.AddSqrt:
        st.Sqrts += n

      case n = <- *st.AddCmp:
        st.Cmps += n

      case ch := <- *st.Stats:
        ch <- st.GetData()
      }
    }
  }()
  wg.Wait()

}

// -------------------------------------------------------------------------------------------------------------------

func (st *ComputeStats) GetData() ComputeStats {
  return ComputeStats{
    Cmps:    st.Cmps,
    Sqrts:   st.Sqrts,
  }
}
