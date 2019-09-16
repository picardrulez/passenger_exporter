package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

//metric definitions
var metricGroupCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "group_count", Help: "Number of groups"})

var metricProcessCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "process_count", Help: "Number of processes"})

var metricMax = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "max", Help: "max"})

var metricCapacityUsed = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "capacity_used", Help: "capacity used"})

var metricGetWaitListSize = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "get_wait_list_size", Help: "size of the wait list that ya get"})

var metricBenefitsPid = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_pid", Help: "PID for Benefits process"})

var metricBenefitsCapUsed = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_capacity_used", Help: "Capacity Used for benefits Process"})

var metricBenefitsEnabledProcCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_enabled_proc_count", Help: "Enabled Proc Count for Benefits"})

var metricBenefitsUID = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_uid", Help: "UID for Benefits Process"})

var metricBenefitsGID = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_gid", Help: "GID for Benefits Process"})

var metricBenefitsStartTimeout = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_start_timeout", Help: "Start Timeout for Benefits Process"})

var metricBenefitsThreadCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_thread_count", Help: "Thread Count for Benefits Process"})

var metricBenefitsMinProcesses = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_min_processes", Help: "Minimum Processes for Benefits Process"})

var metricBenefitsMaxProcesses = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_max_processes", Help: "Maximum Processes for Benefits Process"})

var metricBenefitsStickySessionID = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_sticky_session_id", Help: "Sticky Session ID for Benefits Process"})

var metricBenefitsConcurrency = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_concurrency", Help: "Concurrency for Benefits Process"})

var metricBenefitsSessions = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_sessions", Help: "Number of Benefits Sessions"})

var metricBenefitsProcessed = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_processed", Help: "Total processed in Benefits"})

var metricBenefitsSpawnerCreationTime = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_spawner_creation_time", Help: "Time of spawner creation"})

var metricBenefitsSpawnStartTime = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_spawn_start_time", Help: "Time of spawn start"})

var metricBenefitsSpawnEndTime = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_spawn_end_time", Help: "Time of spawn end"})

var metricBenefitsSwap = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_swap", Help: "Swap Usage for Benefits"})

var metricBenefitsRealMemory = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_real_memory", Help: "Real Memory Usage for Benefits Process"})

var metricBenefitsVMSize = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_vm_size", Help: "VM Size for Benefits"})

var metricBenefitsPGID = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "benefits_pgid", Help: "Process Group ID for Benefits Process"})

func init() {
	//register initial metrics with prometheus
	prometheus.MustRegister(metricGroupCount)
	prometheus.MustRegister(metricProcessCount)
	prometheus.MustRegister(metricMax)
	prometheus.MustRegister(metricCapacityUsed)
	prometheus.MustRegister(metricGetWaitListSize)
	prometheus.MustRegister(metricBenefitsPid)
	prometheus.MustRegister(metricBenefitsCapUsed)
	prometheus.MustRegister(metricBenefitsEnabledProcCount)
	prometheus.MustRegister(metricBenefitsUID)
	prometheus.MustRegister(metricBenefitsGID)
	prometheus.MustRegister(metricBenefitsStartTimeout)
	prometheus.MustRegister(metricBenefitsThreadCount)
	prometheus.MustRegister(metricBenefitsMinProcesses)
	prometheus.MustRegister(metricBenefitsMaxProcesses)
	prometheus.MustRegister(metricBenefitsStickySessionID)
	prometheus.MustRegister(metricBenefitsConcurrency)
	prometheus.MustRegister(metricBenefitsSessions)
	prometheus.MustRegister(metricBenefitsProcessed)
	prometheus.MustRegister(metricBenefitsSpawnerCreationTime)
	prometheus.MustRegister(metricBenefitsSpawnStartTime)
	prometheus.MustRegister(metricBenefitsSpawnEndTime)
	prometheus.MustRegister(metricBenefitsSwap)
	prometheus.MustRegister(metricBenefitsRealMemory)
	prometheus.MustRegister(metricBenefitsVMSize)
	prometheus.MustRegister(metricBenefitsPGID)

	//set values for metrics
	metricGroupCount.Set(0)
	metricProcessCount.Set(0)
	metricMax.Set(0)
	metricCapacityUsed.Set(0)
	metricGetWaitListSize.Set(0)
	metricBenefitsPid.Set(0)
	metricBenefitsCapUsed.Set(0)
	metricBenefitsEnabledProcCount.Set(0)
	metricBenefitsUID.Set(0)
	metricBenefitsGID.Set(0)
	metricBenefitsStartTimeout.Set(0)
	metricBenefitsThreadCount.Set(0)
	metricBenefitsMinProcesses.Set(0)
	metricBenefitsMaxProcesses.Set(0)
	metricBenefitsStickySessionID.Set(0)
	metricBenefitsConcurrency.Set(0)
	metricBenefitsSessions.Set(0)
	metricBenefitsProcessed.Set(0)
	metricBenefitsSpawnerCreationTime.Set(0)
	metricBenefitsSpawnStartTime.Set(0)
	metricBenefitsSpawnEndTime.Set(0)
	metricBenefitsSwap.Set(0)
	metricBenefitsRealMemory.Set(0)
	metricBenefitsVMSize.Set(0)
	metricBenefitsPGID.Set(0)
}

func recordMetrics() {
	go func() {
		for {
			passengerInfo := getXML()
			metricGroupCount.Set(int2float64(passengerInfo.Group_count))
			metricProcessCount.Set(int2float64(passengerInfo.Process_count))
			metricMax.Set(int2float64(passengerInfo.Max))
			metricCapacityUsed.Set(int2float64(passengerInfo.Capacity_used))
			metricGetWaitListSize.Set(int2float64(passengerInfo.Get_wait_list_size))
			metricBenefitsPid.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.PID))
			metricBenefitsCapUsed.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Capacity_used))
			metricBenefitsEnabledProcCount.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Enabled_process_count))
			metricBenefitsUID.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.UID))
			metricBenefitsGID.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.GID))
			metricBenefitsStartTimeout.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Options.Start_timeout))
			metricBenefitsThreadCount.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Options.Thread_count))
			metricBenefitsMinProcesses.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Options.Min_processes))
			metricBenefitsMaxProcesses.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Options.Max_processes))
			metricBenefitsStickySessionID.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Sticky_session_id))
			metricBenefitsConcurrency.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Concurrency))
			metricBenefitsSessions.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Sessions))
			metricBenefitsProcessed.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Processed))
			metricBenefitsSpawnerCreationTime.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Spawner_creation_time))
			metricBenefitsSpawnStartTime.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Spawn_start_time))
			metricBenefitsSpawnEndTime.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Spawn_end_time))
			metricBenefitsSwap.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Swap))
			metricBenefitsRealMemory.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Real_memory))
			metricBenefitsVMSize.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.VMsize))
			metricBenefitsPGID.Set(int2float64(passengerInfo.Supergroups.Supergroup[0].Group.Processes.Process.Process_group_id))
			time.Sleep(1 * time.Second)
		}
	}()
}
