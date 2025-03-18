package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
	"io/ioutil"
	"os"
	"os/exec"
)

type XMLinfo struct {
	Passenger_version  string         `xml:"passenger_version"`
	Group_count        int            `xml:"group_count"`
	Process_count      int            `xml:"process_count"`
	Max                int            `xml:"max"`
	Capacity_used      int            `xml:"capacity_used"`
	Get_wait_list_size int            `xml:"get_wait_list_size"`
	Supergroups        XMLsupergroups `xml:"supergroups"`
}

type XMLsupergroups struct {
	Supergroup []XMLsupergroup `xml:"supergroup"`
}
type XMLsupergroup struct {
	Name               string   `xml:"name"`
	State              string   `xml:"state"`
	Get_wait_list_size int      `xml:"get_wait_list_size"`
	Capacity_used      int      `xml:"capacity_used"`
	Group              XMLgroup `xml:"group"`
}

type XMLgroup struct {
	Default                 bool         `xml:"default,attr"`
	Name                    string       `xml:"name"`
	Component_name          string       `xml:"component_name"`
	App_root                string       `xml:"app_root"`
	App_type                string       `xml:"app_type"`
	Environment             string       `xml:"environment"`
	UUID                    string       `xml:"uuid"`
	Enabled_process_count   int          `xml:"enabled_process_count"`
	Disabling_process_count int          `xml:"disabling_process_count"`
	Disabled_process_count  int          `xml:"disabled_process_count"`
	Capacity_used           int          `xml:"capacity_used"`
	Get_wait_list_size      int          `xml:"get_wait_list_size"`
	Disable_wait_list_size  int          `xml:"disable_wait_list_size"`
	Processes_being_spawned int          `xml:"processes_being_spawned"`
	Life_status             string       `xml:"life_status"`
	User                    string       `xml:"user"`
	UID                     int          `xml:"uid"`
	Group                   string       `xml:"group"`
	GID                     int          `xml:"gid"`
	Options                 XMLoptions   `xml:"options"`
	Processes               XMLprocesses `xml:"processes"`
}
type XMLoptions struct {
	App_root                       string `xml:"app_root"`
	App_group_name                 string `xml:"app_group_name"`
	App_type                       string `xml:"app_type"`
	Start_command                  string `xml:"start_command"`
	Startup_file                   string `xml:"startup_file"`
	Log_level                      string `xml:"log_level"`
	Start_timeout                  int    `xml:"start_timeout"`
	Environment                    string `xml:"environment"`
	Base_uri                       string `xml:"base_uri"`
	Spawn_method                   string `xml:"spawn_method"`
	Default_user                   string `xml:"default_user"`
	Default_group                  string `xml:"default_group"`
	Integration_mode               string `xml:"integration_mode"`
	Ruby                           string `xml:"ruby"`
	Python                         string `xml:"python"`
	Nodejs                         string `xml:"nodejs"`
	Debugger                       bool   `xml:"debugger"`
	Concurrency_model              string `xml:"concurrency_model"`
	Thread_count                   int    `xml:"thread_count"`
	Min_processes                  int    `xml:"min_processes"`
	Max_processes                  int    `xml:"max_processes"`
	Max_preloader_idle_time        int    `xml:"max_preloader_idle_time"`
	Max_out_of_band_work_instances int    `xml:"max_out_of_band_work_instances"`
	Rolling_restart                bool   `xml:"rolling_restart"`
	Ignore_spawn_errors            bool   `xml:"ignore_spawn_errors"`
}

type XMLprocesses struct {
	Process XMLprocess `xml:"process"`
}

type XMLprocess struct {
	PID                   int    `xml:"pid"`
	Sticky_session_id     int    `xml:"sticky_session_id"`
	GUPID                 string `xml:"gupid"`
	Concurrency           int    `xml:"concurrency"`
	Sessions              int    `xml:"sessions"`
	Busyness              int    `xml:"busyness"`
	Processed             int    `xml:"processed"`
	Spawner_creation_time int    `xml:"spawner_creation_time"`
	Spawn_start_time      int    `xml:"spawn_start_time"`
	Spawn_end_time        int    `xml:"spawn_end_time"`
	Last_used             int    `xml:"last_used"`
	Last_used_desc        string `xml:"last_used_desc"`
	Uptime                string `xml:"uptime"`
	Life_status           string `xml:"life_status"`
	Enabled               string `xml:"enabled"`
	Has_metrics           bool   `xml:"has_metrics"`
	CPU                   int    `xml:"cpu"`
	RSS                   int    `xml:"rss"`
	PSS                   int    `xml:"pss"`
	Private_dirty         int    `xml:"private_dirty"`
	Swap                  int    `xml:"swap"`
	Real_memory           int    `xml:"real_memory"`
	VMsize                int    `xml:"vmsize"`
	Process_group_id      int    `xml:"process_group_id"`
	Command               string `xml:"command"`
}

func oldgetXML() {
	xmlFile, err := os.Open("passenger.xml")
	if err != nil {
		log.Info("failed reading file")
		log.Info(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var passengerInfo XMLinfo
	reader := bytes.NewReader(byteValue)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err = decoder.Decode(&passengerInfo)
	if err != nil {
		log.Info("error decoding xml data")
	}
	//	xml.Unmarshal(byteValue, &passengerInfo)
	log.Info("xmlValues:")
	fmt.Printf("%+v", passengerInfo)
	log.Info("printing passenger_version:")
	fmt.Printf(passengerInfo.Passenger_version)
}

func getXML() XMLinfo {
	cmd := exec.Command(EXECCOMMAND, EXECOPTIONS)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("cmd failed with %s\n", err)
	}
	var passengerInfo XMLinfo
	reader := bytes.NewReader([]byte(output))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err = decoder.Decode(&passengerInfo)
	if err != nil {
		log.Info("error decoding xml data")
		log.Info(err)
	}
	//arrayResult := testArray(passengerInfo)
	//if arrayResult == false {
	//		log.Println("array test failed, exiting")
	//		os.Exit(1)
	//	}
	return passengerInfo
}
