package monitors

//Not really a unit test atm
//func TestRPCServer(t *testing.T) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
//	defer cancel()
//	log := logrus.New()
//	monitor, err := GetOsMonitor(log, "linux")
//	go func() {
//		monitor.Run(ctx)
//	}()
//	require.NoError(t, err)
//	require.Len(t, monitor.avgRequired, 0)
//	require.Len(t, monitor.averages, 0)
//
//	monitor.AddMAverage(4)
//	require.Len(t, monitor.avgRequired, 1)
//	require.Len(t, monitor.averages, 0)
//
//	time.Sleep(6 * time.Second)
//	require.Len(t, monitor.averages, 1)
//
//	monitor.AddMAverage(5)
//	time.Sleep(2 * time.Second)
//	require.Len(t, monitor.avgRequired, 2)
//	require.Len(t, monitor.averages, 2)
//
//
//	time.Sleep(1 * time.Second)
//	monitor.AddMAverage(10)
//	require.Len(t, monitor.avgRequired, 3)
//	require.Len(t, monitor.averages, 2)
//
//	time.Sleep(4 * time.Second)
//	require.Len(t, monitor.averages, 3)
//}
