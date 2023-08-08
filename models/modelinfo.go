package models

type ModelInfo struct {
	Custom struct {
		LastWifiChan      int    `json:"lastWifiChan"`
		DsaLocalUrl       string `json:"dsaLocalUrl"`
		HiddenMenuEnabled bool   `json:"hiddenMenuEnabled"`
		HideAdminPassword bool   `json:"hideAdminPassword"`
		CarrierSim        string `json:"carrierSim"`
		End               int    `json:"end"`
	} `json:"custom"`
	Webd struct {
		AdminPassword              string        `json:"adminPassword"`
		OwnerModeEnabled           bool          `json:"ownerModeEnabled"`
		HideAdminPassword          bool          `json:"hideAdminPassword"`
		HintAnswer                 string        `json:"hintAnswer"`
		HintNumber                 int           `json:"hintNumber"`
		OwnerCustomizationsChanged bool          `json:"ownerCustomizationsChanged"`
		OwnerCustomizationsList    []interface{} `json:"ownerCustomizationsList"`
		End                        string        `json:"end"`
	} `json:"webd"`
	Lcd struct {
		BacklightEnabled       bool   `json:"backlightEnabled"`
		BacklightActive        bool   `json:"backlightActive"`
		InactivityTimer        int    `json:"inactivityTimer"`
		InactivityTimerAC      int    `json:"inactivityTimerAC"`
		InactivityTimerUSB     int    `json:"inactivityTimerUSB"`
		BrightnessOnBatt       string `json:"brightnessOnBatt"`
		BrightnessOnUSB        string `json:"brightnessOnUSB"`
		BrightnessOnAC         string `json:"brightnessOnAC"`
		BrightnessOnBattIntVal int    `json:"brightnessOnBattIntVal"`
		BrightnessOnUSBIntVal  int    `json:"brightnessOnUSBIntVal"`
		BrightnessOnACIntVal   int    `json:"brightnessOnACIntVal"`
		BacklightOverride      string `json:"backlightOverride"`
		LockScreenEnable       bool   `json:"lockScreenEnable"`
		LockScreenUsePin       bool   `json:"lockScreenUsePin"`
		End                    string `json:"end"`
	} `json:"lcd"`
	Led struct {
		End string `json:"end"`
	} `json:"led"`
	Sim struct {
		Pin struct {
			Mode  string `json:"mode"`
			Retry int    `json:"retry"`
			End   string `json:"end"`
		} `json:"pin"`
		Puk struct {
			Retry int `json:"retry"`
		} `json:"puk"`
		Mep struct {
		} `json:"mep"`
		PhoneNumber   string `json:"phoneNumber"`
		Iccid         string `json:"iccid"`
		Imsi          string `json:"imsi"`
		SPN           string `json:"SPN"`
		Status        string `json:"status"`
		SprintSimLock int    `json:"sprintSimLock"`
		End           string `json:"end"`
	} `json:"sim"`
	Sms struct {
		Ready          bool   `json:"ready"`
		SendSupported  bool   `json:"sendSupported"`
		SendEnabled    bool   `json:"sendEnabled"`
		UnreadMsgs     int    `json:"unreadMsgs"`
		AlertSupported bool   `json:"alertSupported"`
		AlertEnabled   bool   `json:"alertEnabled"`
		AlertNumList   string `json:"alertNumList"`
		MsgCount       int    `json:"msgCount"`
		Msgs           []struct {
			Id     string `json:"id,omitempty"`
			RxTime string `json:"rxTime,omitempty"`
			Text   string `json:"text,omitempty"`
			Sender string `json:"sender,omitempty"`
			Read   bool   `json:"read,omitempty"`
		} `json:"msgs"`
		Trans   []interface{} `json:"trans"`
		SendMsg []struct {
			ClientId  string `json:"clientId,omitempty"`
			Enc       string `json:"enc,omitempty"`
			ErrorCode int    `json:"errorCode,omitempty"`
			MsgId     int    `json:"msgId,omitempty"`
			Receiver  string `json:"receiver,omitempty"`
			Status    string `json:"status,omitempty"`
			Text      string `json:"text,omitempty"`
			TxTime    string `json:"txTime,omitempty"`
		} `json:"sendMsg"`
		End string `json:"end"`
	} `json:"sms"`
	Session struct {
		UserRole            string `json:"userRole"`
		Lang                string `json:"lang"`
		HintDisplayPassword string `json:"hintDisplayPassword"`
		SecToken            string `json:"secToken"`
		ClientIP            string `json:"clientIP"`
		SupportedLangList   []struct {
			Id        string `json:"id,omitempty"`
			IsCurrent bool   `json:"isCurrent,omitempty"`
			Label     string `json:"label,omitempty"`
		} `json:"supportedLangList"`
	} `json:"session"`
	General struct {
		DefaultLanguage           string        `json:"defaultLanguage"`
		PRIid                     string        `json:"PRIid"`
		ActivationDate            string        `json:"activationDate"`
		GenericResetStatus        string        `json:"genericResetStatus"`
		ReconditionDate           string        `json:"reconditionDate"`
		Manufacturer              string        `json:"manufacturer"`
		Model                     string        `json:"model"`
		HWversion                 string        `json:"HWversion"`
		FWversion                 string        `json:"FWversion"`
		ModemFwVersion            string        `json:"modemFwVersion"`
		BuildDate                 string        `json:"buildDate"`
		BLversion                 string        `json:"BLversion"`
		PRIversion                string        `json:"PRIversion"`
		TruInstallAvailable       bool          `json:"truInstallAvailable"`
		TruInstallVersion         string        `json:"truInstallVersion"`
		IMEI                      string        `json:"IMEI"`
		SVN                       string        `json:"SVN"`
		MEID                      string        `json:"MEID"`
		ESN                       string        `json:"ESN"`
		FSN                       string        `json:"FSN"`
		ATTDeviceId               string        `json:"ATTDeviceId"`
		PackageName               string        `json:"packageName"`
		Activated                 bool          `json:"activated"`
		WebAppVersion             string        `json:"webAppVersion"`
		AppVersion                string        `json:"appVersion"`
		HIDenabled                bool          `json:"HIDenabled"`
		TruInstallEnabled         bool          `json:"truInstallEnabled"`
		TruInstallSupported       bool          `json:"truInstallSupported"`
		TCAaccepted               bool          `json:"TCAaccepted"`
		LEDenabled                bool          `json:"LEDenabled"`
		ShowAdvHelp               bool          `json:"showAdvHelp"`
		ShowWebInfo               bool          `json:"showWebInfo"`
		KeyLockState              string        `json:"keyLockState"`
		DevTemperature            int           `json:"devTemperature"`
		VerMajor                  int           `json:"verMajor"`
		VerMinor                  int           `json:"verMinor"`
		Environment               string        `json:"environment"`
		CurrTime                  int           `json:"currTime"`
		TimeZoneOffset            int           `json:"timeZoneOffset"`
		UserTzOffsetEnabled       bool          `json:"userTzOffsetEnabled"`
		UserTzOffset              int           `json:"userTzOffset"`
		DeviceName                string        `json:"deviceName"`
		UseMetricSystem           bool          `json:"useMetricSystem"`
		FactoryResetButtonEnabled bool          `json:"factoryResetButtonEnabled"`
		FactoryResetLCDSupported  bool          `json:"factoryResetLCDSupported"`
		FactoryResetStatus        string        `json:"factoryResetStatus"`
		SPClockStatus             string        `json:"SPClockStatus"`
		SetupCompleted            bool          `json:"setupCompleted"`
		WarrantyDateCode          string        `json:"warrantyDateCode"`
		LanguageSelected          bool          `json:"languageSelected"`
		UiDateFormat              string        `json:"uiDateFormat"`
		UpTime                    int           `json:"upTime"`
		Use24HrTimeFormat         bool          `json:"use24HrTimeFormat"`
		SystemAlertList           []interface{} `json:"systemAlertList"`
		ApiVersion                string        `json:"apiVersion"`
		CompanyName               string        `json:"companyName"`
		UsbDevicesURL             string        `json:"usbDevicesURL"`
		NwFoldersURL              string        `json:"nwFoldersURL"`
		ConfigURL                 string        `json:"configURL"`
		ProfileURL                string        `json:"profileURL"`
		PinChangeURL              string        `json:"pinChangeURL"`
		PortCfgURL                string        `json:"portCfgURL"`
		PortFilterURL             string        `json:"portFilterURL"`
		WifiACLURL                string        `json:"wifiACLURL"`
		SupportedLangList         []struct {
			Id        string `json:"id,omitempty"`
			IsCurrent string `json:"isCurrent,omitempty"`
			IsDefault string `json:"isDefault,omitempty"`
			Label     string `json:"label,omitempty"`
			Token1    string `json:"token1,omitempty"`
			Token2    string `json:"token2,omitempty"`
		} `json:"supportedLangList"`
	} `json:"general"`
	Power struct {
		PMState          string `json:"PMState"`
		SmState          string `json:"SmState"`
		BattLowThreshold int    `json:"battLowThreshold"`
		AutoOff          struct {
			OnUSBdisconnect struct {
				Enable         bool   `json:"enable"`
				CountdownTimer int    `json:"countdownTimer"`
				End            string `json:"end"`
			} `json:"onUSBdisconnect"`
			OnIdle struct {
				Timer struct {
					OnAC      int    `json:"onAC"`
					OnBattery int    `json:"onBattery"`
					End       string `json:"end"`
				} `json:"timer"`
			} `json:"onIdle"`
		} `json:"autoOff"`
		Standby struct {
			Ethernet struct {
				Timer struct {
					OnBattery int    `json:"onBattery"`
					End       string `json:"end"`
				} `json:"timer"`
			} `json:"ethernet"`
			OnIdle struct {
				Timer struct {
					OnAC      int    `json:"onAC"`
					OnBattery int    `json:"onBattery"`
					OnUSB     int    `json:"onUSB"`
					End       string `json:"end"`
				} `json:"timer"`
			} `json:"onIdle"`
		} `json:"standby"`
		AutoOn struct {
			Enable bool   `json:"enable"`
			End    string `json:"end"`
		} `json:"autoOn"`
		BatteryTemperature  int    `json:"batteryTemperature"`
		BatteryVoltage      int    `json:"batteryVoltage"`
		BattChargeLevel     int    `json:"battChargeLevel"`
		BattChargeSource    string `json:"battChargeSource"`
		BatteryState        string `json:"batteryState"`
		BattChargeAlgorithm string `json:"battChargeAlgorithm"`
		Charging            bool   `json:"charging"`
		ButtonHoldTime      int    `json:"buttonHoldTime"`
		DeviceTempCritical  bool   `json:"deviceTempCritical"`
		Resetreason         int    `json:"resetreason"`
		ResetRequired       string `json:"resetRequired"`
		ChoiceOnUsb         string `json:"choiceOnUsb"`
		ActionOnUsb         string `json:"actionOnUsb"`
		BatteryTempState    string `json:"batteryTempState"`
		BatteryName         string `json:"batteryName"`
		WifiOff             struct {
			OnUsbConnect bool   `json:"onUsbConnect"`
			OnTethered   bool   `json:"onTethered"`
			End          string `json:"end"`
		} `json:"wifiOff"`
		Boost struct {
			CableConnected bool   `json:"cableConnected"`
			End            string `json:"end"`
		} `json:"boost"`
		Lpm bool   `json:"lpm"`
		End string `json:"end"`
	} `json:"power"`
	Wwan struct {
		PrlVersion               int    `json:"prlVersion"`
		BandDisablementMaskLTE   string `json:"bandDisablementMaskLTE"`
		BandDisablementMask      string `json:"bandDisablementMask"`
		LTEBandPriority          string `json:"LTEBandPriority"`
		NetScanStatus            string `json:"netScanStatus"`
		MtuSize                  int    `json:"mtuSize"`
		MtuChangeEnabled         bool   `json:"mtuChangeEnabled"`
		LTEeHRPDConfig           string `json:"LTEeHRPDConfig"`
		RoamingEnhancedIndicator int    `json:"roamingEnhancedIndicator"`
		RoamingMode              string `json:"roamingMode"`
		RoamingGuardDom          string `json:"roamingGuardDom"`
		RoamingGuardIntl         string `json:"roamingGuardIntl"`
		RoamingType              string `json:"roamingType"`
		RoamMenuDisplay          bool   `json:"roamMenuDisplay"`
		ERITestMode              bool   `json:"ERITestMode"`
		AutoBandRegionChanged    bool   `json:"autoBandRegionChanged"`
		InactivityCause          int    `json:"inactivityCause"`
		CurrentNWserviceType     string `json:"currentNWserviceType"`
		RegisterRejectCode       int    `json:"registerRejectCode"`
		NetSelEnabled            string `json:"netSelEnabled"`
		NetRegMode               string `json:"netRegMode"`
		Roaming                  bool   `json:"roaming"`
		IPv6                     string `json:"IPv6"`
		IP                       string `json:"IP"`
		RegisterNetworkDisplay   string `json:"registerNetworkDisplay"`
		RAT                      string `json:"RAT"`
		BandRegion               []struct {
			Index   int    `json:"index,omitempty"`
			Name    string `json:"name,omitempty"`
			Current bool   `json:"current,omitempty"`
		} `json:"bandRegion"`
		Autoconnect string `json:"autoconnect"`
		ProfileList []struct {
			Index          int    `json:"index,omitempty"`
			Id             string `json:"id,omitempty"`
			Name           string `json:"name,omitempty"`
			Apn            string `json:"apn,omitempty"`
			Username       string `json:"username,omitempty"`
			Password       string `json:"password,omitempty"`
			Authtype       string `json:"authtype,omitempty"`
			Ipaddr         string `json:"ipaddr,omitempty"`
			AccessControl  int    `json:"access_control,omitempty"`
			Type           string `json:"type,omitempty"`
			Pdproamingtype string `json:"pdproamingtype,omitempty"`
		} `json:"profileList"`
		Profile struct {
			Default               string `json:"default"`
			DefaultLTE            string `json:"defaultLTE"`
			PromptForApnSelection bool   `json:"promptForApnSelection"`
			End                   string `json:"end"`
		} `json:"profile"`
		PromptForPwd string `json:"promptForPwd"`
		DataUsage    struct {
			Total struct {
				LteBillingTx  int    `json:"lteBillingTx"`
				LteBillingRx  int    `json:"lteBillingRx"`
				CdmaBillingTx int    `json:"cdmaBillingTx"`
				CdmaBillingRx int    `json:"cdmaBillingRx"`
				GwBillingTx   int    `json:"gwBillingTx"`
				GwBillingRx   int    `json:"gwBillingRx"`
				LteLifeTx     int    `json:"lteLifeTx"`
				LteLifeRx     int    `json:"lteLifeRx"`
				CdmaLifeTx    int    `json:"cdmaLifeTx"`
				CdmaLifeRx    int    `json:"cdmaLifeRx"`
				GwLifeTx      int    `json:"gwLifeTx"`
				GwLifeRx      int    `json:"gwLifeRx"`
				End           string `json:"end"`
			} `json:"total"`
			Server struct {
				AccountType    string `json:"accountType"`
				SubAccountType string `json:"subAccountType"`
				End            string `json:"end"`
			} `json:"server"`
			Remote struct {
				Enabled bool   `json:"enabled"`
				End     string `json:"end"`
			} `json:"remote"`
			ServerDataRemaining       int    `json:"serverDataRemaining"`
			ServerDataTransferred     int    `json:"serverDataTransferred"`
			ServerDataTransferredIntl int    `json:"serverDataTransferredIntl"`
			ServerDataValidState      string `json:"serverDataValidState"`
			ServerDaysLeft            int    `json:"serverDaysLeft"`
			ServerErrorCode           string `json:"serverErrorCode"`
			ServerLowBalance          bool   `json:"serverLowBalance"`
			ServerMsisdn              string `json:"serverMsisdn"`
			ServerRechargeUrl         string `json:"serverRechargeUrl"`
			DataWarnEnable            bool   `json:"dataWarnEnable"`
			PlanSize                  int    `json:"planSize"`
			PlanDescription           string `json:"planDescription"`
			PrepaidStackedPlans       int    `json:"prepaidStackedPlans"`
			PrepaidStackedPlansIntl   int    `json:"prepaidStackedPlansIntl"`
			PrepaidAccountState       string `json:"prepaidAccountState"`
			AccountType               string `json:"accountType"`
			DisableAutoReset          bool   `json:"disableAutoReset"`
			Share                     struct {
				Enabled               bool   `json:"enabled"`
				DataTransferredOthers int    `json:"dataTransferredOthers"`
				LastSync              string `json:"lastSync"`
				End                   string `json:"end"`
			} `json:"share"`
			Generic struct {
				DataLimitValid         bool   `json:"dataLimitValid"`
				UsageHighWarning       int    `json:"usageHighWarning"`
				FallbackSupported      bool   `json:"fallbackSupported"`
				LastSucceeded          string `json:"lastSucceeded"`
				BillingDay             int    `json:"billingDay"`
				NextBillingDate        string `json:"nextBillingDate"`
				LastSync               string `json:"lastSync"`
				BillingCycleRemainder  string `json:"billingCycleRemainder"`
				BillingCycleLimit      int    `json:"billingCycleLimit"`
				DataTransferred        int    `json:"dataTransferred"`
				DataTransferredRoaming int    `json:"dataTransferredRoaming"`
				LastReset              string `json:"lastReset"`
				UserDisplayFormat      string `json:"userDisplayFormat"`
				End                    string `json:"end"`
			} `json:"generic"`
		} `json:"dataUsage"`
		DataTransferred      string `json:"dataTransferred"`
		DataTransferredRx    string `json:"dataTransferredRx"`
		DataTransferredTx    string `json:"dataTransferredTx"`
		NetManualNoCvg       bool   `json:"netManualNoCvg"`
		Connection           string `json:"connection"`
		ConnectionType       string `json:"connectionType"`
		CurrentPSserviceType string `json:"currentPSserviceType"`
		Ca                   struct {
			SCCcount int `json:"SCCcount"`
			SCClist  []struct {
				Cellid  string `json:"cellid,omitempty"`
				Dlchan  string `json:"dlchan,omitempty"`
				Sigrsrp string `json:"sigrsrp,omitempty"`
				Sigrsrq string `json:"sigrsrq,omitempty"`
			} `json:"SCClist"`
			End string `json:"end"`
		} `json:"ca"`
		ConnectionText string `json:"connectionText"`
		SessDuration   int    `json:"sessDuration"`
		SessStartTime  int    `json:"sessStartTime"`
		SignalStrength struct {
			Rssi int    `json:"rssi"`
			Rscp int    `json:"rscp"`
			Ecio int    `json:"ecio"`
			Rsrp int    `json:"rsrp"`
			Rsrq int    `json:"rsrq"`
			Bars int    `json:"bars"`
			Sinr int    `json:"sinr"`
			End  string `json:"end"`
		} `json:"signalStrength"`
		DiagInfo []struct {
			LteAttached       bool   `json:"lteAttached,omitempty"`
			Nr5GAttached      bool   `json:"nr5gAttached,omitempty"`
			EndcEnabledConfig bool   `json:"endcEnabledConfig,omitempty"`
			LtesigValid       bool   `json:"ltesigValid,omitempty"`
			LtesigRssi        string `json:"ltesigRssi,omitempty"`
			LtesigRsrp        string `json:"ltesigRsrp,omitempty"`
			LtesigRsrq        string `json:"ltesigRsrq,omitempty"`
			LtesigSnr         string `json:"ltesigSnr,omitempty"`
			Nr5GsigValid      bool   `json:"nr5gsigValid,omitempty"`
			Nr5GsigRsrp       string `json:"nr5gsigRsrp,omitempty"`
			Nr5GsigRsrq       string `json:"nr5gsigRsrq,omitempty"`
			Nr5GsigSnr        string `json:"nr5gsigSnr,omitempty"`
		} `json:"diagInfo"`
		LteBandInfo []struct {
			IsPcc          bool   `json:"isPcc,omitempty"`
			SccId          int    `json:"sccId,omitempty"`
			SccState       int    `json:"sccState,omitempty"`
			Band           int    `json:"band,omitempty"`
			Channel        string `json:"channel,omitempty"`
			PhyCid         string `json:"phyCid,omitempty"`
			DlBandwidth    string `json:"dlBandwidth,omitempty"`
			SccUlCaEnabled string `json:"sccUlCaEnabled,omitempty"`
			SccUlChannel   int    `json:"sccUlChannel,omitempty"`
			SccUlBandwidth string `json:"sccUlBandwidth,omitempty"`
		} `json:"lteBandInfo"`
		Laa struct {
			Supported bool   `json:"supported"`
			Enabled   bool   `json:"enabled"`
			End       string `json:"end"`
		} `json:"laa"`
	} `json:"wwan"`
	Wwanadv struct {
		CurBand           string `json:"curBand"`
		RadioQuality      int    `json:"radioQuality"`
		Country           string `json:"country"`
		RAC               int    `json:"RAC"`
		LAC               int    `json:"LAC"`
		MCC               string `json:"MCC"`
		MNC               string `json:"MNC"`
		MNCFmt            int    `json:"MNCFmt"`
		CellId            int    `json:"cellId"`
		ChanId            int    `json:"chanId"`
		PrimScode         int    `json:"primScode"`
		PlmnSrvErrBitMask int    `json:"plmnSrvErrBitMask"`
		ChanIdUl          int    `json:"chanIdUl"`
		TxLevel           int    `json:"txLevel"`
		RxLevel           int    `json:"rxLevel"`
		End               string `json:"end"`
	} `json:"wwanadv"`
	Ethernet struct {
		Mac     string `json:"mac"`
		Offload struct {
			Supported bool   `json:"supported"`
			Enabled   bool   `json:"enabled"`
			On        bool   `json:"on"`
			Ipv4Addr  string `json:"ipv4Addr"`
			Ipv6Addr  string `json:"ipv6Addr"`
			Rx        int    `json:"rx"`
			Tx        int    `json:"tx"`
			TimeOn    int    `json:"timeOn"`
			End       string `json:"end"`
		} `json:"offload"`
	} `json:"ethernet"`
	Wifi struct {
		Supported          string `json:"supported"`
		Enabled            bool   `json:"enabled"`
		Status             string `json:"status"`
		Mode               string `json:"mode"`
		MaxClientSupported int    `json:"maxClientSupported"`
		MaxClientLimit     int    `json:"maxClientLimit"`
		MaxClientCnt       int    `json:"maxClientCnt"`
		Encryption         string `json:"encryption"`
		Channel            int    `json:"channel"`
		GBandwidth         string `json:"2gBandwidth"`
		GBandwidth1        string `json:"5gBandwidth"`
		TxPower            string `json:"txPower"`
		HiddenSSID         bool   `json:"hiddenSSID"`
		PassPhrase         string `json:"passPhrase"`
		PasswordReminder   bool   `json:"passwordReminder"`
		RTSthreshold       int    `json:"RTSthreshold"`
		FragThreshold      int    `json:"fragThreshold"`
		AccessControl      string `json:"accessControl"`
		SSID               string `json:"SSID"`
		WmmEnabled         bool   `json:"wmmEnabled"`
		MAC                string `json:"MAC"`
		SSIDreminder       bool   `json:"SSIDreminder"`
		ClientCount        int    `json:"clientCount"`
		Country            string `json:"country"`
		Privacy            bool   `json:"privacy"`
		GMcsChanList       string `json:"2gMcsChanList"`
		GMcsChanList1      string `json:"5gMcsChanList"`
		Wps                struct {
			Supported string `json:"supported"`
			Enabled   string `json:"enabled"`
			Blocked   bool   `json:"blocked"`
			Mode      string `json:"mode"`
			Status    string `json:"status"`
			End       string `json:"end"`
		} `json:"wps"`
		Guest struct {
			MaxClientCnt       int    `json:"maxClientCnt"`
			Enabled            bool   `json:"enabled"`
			Status             string `json:"status"`
			Encryption         string `json:"encryption"`
			SSID               string `json:"SSID"`
			PassPhrase         string `json:"passPhrase"`
			GeneratePassphrase bool   `json:"generatePassphrase"`
			AccessProfile      string `json:"accessProfile"`
			HiddenSSID         bool   `json:"hiddenSSID"`
			TimerEnable        bool   `json:"timerEnable"`
			TimerTimestamp     int    `json:"timerTimestamp"`
			TimerValue         int    `json:"timerValue"`
			Privacy            bool   `json:"privacy"`
			Chan               int    `json:"chan"`
			Mode               string `json:"mode"`
			DHCP               struct {
				Range struct {
					High string `json:"high"`
					Low  string `json:"low"`
					End  string `json:"end"`
				} `json:"range"`
			} `json:"DHCP"`
		} `json:"guest"`
		Aux struct {
			Enabled      bool   `json:"enabled"`
			AuxMode      string `json:"AuxMode"`
			SSID         string `json:"SSID"`
			HiddenSSID   bool   `json:"hiddenSSID"`
			Encryption   string `json:"encryption"`
			PassPhrase   string `json:"passPhrase"`
			MaxClientCnt int    `json:"maxClientCnt"`
			Status       string `json:"status"`
			Chan         int    `json:"chan"`
			Mode         string `json:"mode"`
			End          string `json:"end"`
		} `json:"aux"`
		Offload struct {
			ActivationRequired bool   `json:"activationRequired"`
			Bars               int    `json:"bars"`
			Enabled            bool   `json:"enabled"`
			Rssi               int    `json:"rssi"`
			SecurityStatus     string `json:"securityStatus"`
			Status             string `json:"status"`
			Supported          bool   `json:"supported"`
			ConnectionSsid     string `json:"connectionSsid"`
			ScanProgress       string `json:"scanProgress"`
			StationIPv4        string `json:"stationIPv4"`
			TimeOn             int    `json:"timeOn"`
			DataTransferred    struct {
				Rx  int    `json:"rx"`
				Tx  int    `json:"tx"`
				End string `json:"end"`
			} `json:"dataTransferred"`
			NetworkList []interface{} `json:"networkList"`
			ScanList    []interface{} `json:"scanList"`
			End         string        `json:"end"`
		} `json:"offload"`
		AccessBlackList []interface{} `json:"accessBlackList"`
		AccessWhiteList []interface{} `json:"accessWhiteList"`
		End             string        `json:"end"`
	} `json:"wifi"`
	Router struct {
		GatewayIP  string `json:"gatewayIP"`
		DMZaddress string `json:"DMZaddress"`
		DMZenabled bool   `json:"DMZenabled"`
		ForceSetup bool   `json:"forceSetup"`
		DHCP       struct {
			ServerEnabled bool   `json:"serverEnabled"`
			DNS1          string `json:"DNS1"`
			DNS2          string `json:"DNS2"`
			DNSmode       string `json:"DNSmode"`
			USBpcIP       string `json:"USBpcIP"`
			Range         struct {
				High string `json:"high"`
				Low  string `json:"low"`
				End  string `json:"end"`
			} `json:"range"`
		} `json:"DHCP"`
		UsbMode              string        `json:"usbMode"`
		UsbNetworkTethering  bool          `json:"usbNetworkTethering"`
		PortFwdEnabled       bool          `json:"portFwdEnabled"`
		UsbTetheringActive   bool          `json:"usbTetheringActive"`
		PortFwdList          []interface{} `json:"portFwdList"`
		PortFwdAllowEntry    int           `json:"portFwdAllowEntry"`
		PortFilteringEnabled bool          `json:"portFilteringEnabled"`
		PortFilteringMode    string        `json:"portFilteringMode"`
		PortFilterWhiteList  []interface{} `json:"portFilterWhiteList"`
		PortFilterBlackList  []interface{} `json:"portFilterBlackList"`
		ClientList           []struct {
			IP     string `json:"IP,omitempty"`
			MAC    string `json:"MAC,omitempty"`
			Name   string `json:"name,omitempty"`
			OnUSB  bool   `json:"onUSB,omitempty"`
			Source string `json:"source,omitempty"`
		} `json:"clientList"`
		HostName               string `json:"hostName"`
		DomainName             string `json:"domainName"`
		IpPassThroughEnabled   bool   `json:"ipPassThroughEnabled"`
		IpPassThroughSupported bool   `json:"ipPassThroughSupported"`
		VPNpassthrough         bool   `json:"VPNpassthrough"`
		UsbTetherSupported     bool   `json:"usbTetherSupported"`
		NetMask                string `json:"netMask"`
		BridgeLanNetMask       string `json:"bridgeLanNetMask"`
		Ipv6Supported          bool   `json:"Ipv6Supported"`
		UPNPsupported          bool   `json:"UPNPsupported"`
		UPNPenabled            bool   `json:"UPNPenabled"`
		UPNPIgdVersion         int    `json:"UPNPIgdVersion"`
		End                    string `json:"end"`
	} `json:"router"`
	Fota struct {
		Fwupdater struct {
			Available     bool   `json:"available"`
			Chkallow      bool   `json:"chkallow"`
			Chkstatus     string `json:"chkstatus"`
			DloadProg     int    `json:"dloadProg"`
			Error         bool   `json:"error"`
			LastChkDate   int    `json:"lastChkDate"`
			State         string `json:"state"`
			IsPostponable bool   `json:"isPostponable"`
			StatusCode    int    `json:"statusCode"`
			ChkTimeLeft   int    `json:"chkTimeLeft"`
			DloadSize     int    `json:"dloadSize"`
			IsRejectable  bool   `json:"isRejectable"`
			Description   string `json:"description"`
			End           string `json:"end"`
		} `json:"fwupdater"`
	} `json:"fota"`
	Failover struct {
		Enabled         bool   `json:"enabled"`
		Backhaul        string `json:"backhaul"`
		Supported       bool   `json:"supported"`
		MonitorPeriod   int    `json:"monitorPeriod"`
		WanConnected    bool   `json:"wanConnected"`
		KeepaliveEnable bool   `json:"keepaliveEnable"`
		KeepaliveSleep  int    `json:"keepaliveSleep"`
		Ipv4Targets     []struct {
			Id     string `json:"id,omitempty"`
			String string `json:"string,omitempty"`
		} `json:"ipv4Targets"`
		Ipv6Targets []interface{} `json:"ipv6Targets"`
		End         string        `json:"end"`
	} `json:"failover"`
	Cradle struct {
		Mode                bool   `json:"mode"`
		SmartMode           bool   `json:"smartMode"`
		Url                 string `json:"url"`
		PrimarySSID         string `json:"primarySSID"`
		PrimaryPassphrase   string `json:"primaryPassphrase"`
		SecondarySSID       string `json:"secondarySSID"`
		SecondaryPassphrase string `json:"secondaryPassphrase"`
		AutoIPT             bool   `json:"autoIPT"`
		End                 string `json:"end"`
	} `json:"cradle"`
	Eventlog struct {
		Level       int `json:"level"`
		LogDuration int `json:"logDuration"`
		CatLevel    []struct {
			Category     string `json:"category,omitempty"`
			CatlevelMask int    `json:"catlevel_mask,omitempty"`
			Enabled      bool   `json:"enabled,omitempty"`
		} `json:"catLevel"`
		End int `json:"end"`
	} `json:"eventlog"`
	Ui struct {
		ServerDaysLeftHide   bool `json:"serverDaysLeftHide"`
		PromptActivation     bool `json:"promptActivation"`
		StealthEnabled       bool `json:"stealthEnabled"`
		DisableSettingsOnLCD bool `json:"disableSettingsOnLCD"`
		End                  int  `json:"end"`
	} `json:"ui"`
	Accesscontrol struct {
		Nlpc struct {
			FilterLevel string        `json:"filterLevel"`
			Enabled     bool          `json:"enabled"`
			Username    string        `json:"username"`
			MacList     []interface{} `json:"macList"`
			End         int           `json:"end"`
		} `json:"nlpc"`
		Blocksites struct {
			Enabled     bool          `json:"enabled"`
			MacMode     string        `json:"macMode"`
			MacList     []interface{} `json:"macList"`
			PatternMode string        `json:"patternMode"`
			PatternList []interface{} `json:"patternList"`
			End         int           `json:"end"`
		} `json:"blocksites"`
		SchedulerBlockingActive bool   `json:"schedulerBlockingActive"`
		SchedulerSupported      bool   `json:"schedulerSupported"`
		SchedulerEnabled        bool   `json:"schedulerEnabled"`
		SchedulerDefaultMode    string `json:"schedulerDefaultMode"`
		SchedulerList           []struct {
		} `json:"schedulerList"`
		End int `json:"end"`
	} `json:"accesscontrol"`
	Ready struct {
		ShareSupported bool `json:"shareSupported"`
		ShareEnabled   bool `json:"shareEnabled"`
		CloudSupported bool `json:"cloudSupported"`
		CloudEnabled   bool `json:"cloudEnabled"`
		DeviceShare    struct {
			NetworkDeviceName     string        `json:"networkDeviceName"`
			WorkgroupName         string        `json:"workgroupName"`
			RemoveUsbDeviceResult string        `json:"removeUsbDeviceResult"`
			ExtMediaSupported     bool          `json:"extMediaSupported"`
			UsbDevicesInfo        []interface{} `json:"usbDevicesInfo"`
			NwFoldersInfo         []interface{} `json:"nwFoldersInfo"`
			End                   int           `json:"end"`
		} `json:"deviceShare"`
		Cloud struct {
			RegistrationStatus string `json:"registrationStatus"`
			End                int    `json:"end"`
		} `json:"cloud"`
	} `json:"ready"`
	Insight struct {
		Supported bool   `json:"supported"`
		Enabled   bool   `json:"enabled"`
		Active    bool   `json:"active"`
		Version   string `json:"version"`
		End       string `json:"end"`
	} `json:"insight"`
}
