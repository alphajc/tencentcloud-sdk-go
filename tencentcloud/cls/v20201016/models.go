// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201016

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlarmInfo struct {

	// 告警策略名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 触发条件。
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是min。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知模板列表。
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// 开启状态。
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 自定义通知模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// 自定义回调模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`

	// 多维分析设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

type AlarmNotice struct {

	// 告警通知模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发
	// <br><li> Recovery - 告警恢复
	// <br><li> All - 告警触发和告警恢复
	Type *string `json:"Type,omitempty" name:"Type"`

	// 告警通知模板接收者信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// 告警通知模板回调信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`

	// 告警通知模板ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AlarmTarget struct {

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 查询语句。
	Query *string `json:"Query,omitempty" name:"Query"`

	// 告警对象序号；从1开始递增。
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 日志集ID。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

type AlarmTargetInfo struct {

	// 日志集ID。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集名称。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 查询语句。
	Query *string `json:"Query,omitempty" name:"Query"`

	// 告警对象序号。
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type AnalysisDimensional struct {

	// 分析名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分析类型：query，field
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分析内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ApplyConfigToMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 采集配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ApplyConfigToMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConfigToMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyConfigToMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyConfigToMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallBackInfo struct {

	// 回调时的Body
	Body *string `json:"Body,omitempty" name:"Body"`

	// 回调时的Headers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*string `json:"Headers,omitempty" name:"Headers"`
}

type Ckafka struct {

	// Ckafka 的 Vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Ckafka 的 Vport
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// Ckafka 的 InstanceId
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Ckafka 的 InstanceName
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Ckafka 的 TopicId
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Ckafka 的 TopicName
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CloseKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS对应的topic标识
	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {

	// 列的名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 列的属性
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CompressInfo struct {

	// 压缩格式，支持gzip、lzop和none不压缩
	Format *string `json:"Format,omitempty" name:"Format"`
}

type ConfigExtraInfo struct {

	// 采集规则扩展配置ID
	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`

	// 采集规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 类型：container_stdout、container_file、host_file
	Type *string `json:"Type,omitempty" name:"Type"`

	// 节点文件配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`

	// 容器文件路径信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`

	// 容器标准输出信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`

	// 日志格式化方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户自定义解析字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 自建采集配置标
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`

	// 日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集name
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题name
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ConfigInfo struct {

	// 采集规则配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 日志格式化方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`

	// 日志采集路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// 采集配置所属日志主题ID即TopicId
	Output *string `json:"Output,omitempty" name:"Output"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户自定义解析字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

type ConsumerContent struct {

	// 是否投递 TAG 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`

	// 需要投递的元数据列表，目前仅支持：__SOURCE__，__FILENAME__和__TIMESTAMP__
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`

	// 当EnableTag为true时，必须填写TagJsonNotTiled字段，TagJsonNotTiled用于标识tag信息是否json平铺，TagJsonNotTiled为true时不平铺，false时平铺
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagJsonNotTiled *bool `json:"TagJsonNotTiled,omitempty" name:"TagJsonNotTiled"`
}

type ContainerFileInfo struct {

	// namespace可以多个，用分隔号分割,例如A,B
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 容器名称
	Container *string `json:"Container,omitempty" name:"Container"`

	// 日志文件夹
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// 日志名称
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeLabels []*string `json:"IncludeLabels,omitempty" name:"IncludeLabels"`

	// 工作负载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkLoad *ContainerWorkLoadInfo `json:"WorkLoad,omitempty" name:"WorkLoad"`

	// 需要排除的namespace可以多个，用分隔号分割,例如A,B
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeNamespace *string `json:"ExcludeNamespace,omitempty" name:"ExcludeNamespace"`

	// 需要排除的pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeLabels []*string `json:"ExcludeLabels,omitempty" name:"ExcludeLabels"`
}

type ContainerStdoutInfo struct {

	// 是否所有容器
	AllContainers *bool `json:"AllContainers,omitempty" name:"AllContainers"`

	// container为空表所有的，不为空采集指定的容器
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitempty" name:"Container"`

	// namespace可以多个，用分隔号分割,例如A,B；为空或者没有这个字段，表示所有namespace
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeLabels []*string `json:"IncludeLabels,omitempty" name:"IncludeLabels"`

	// 工作负载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkLoads []*ContainerWorkLoadInfo `json:"WorkLoads,omitempty" name:"WorkLoads"`

	// 需要排除的namespace可以多个，用分隔号分割,例如A,B
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeNamespace *string `json:"ExcludeNamespace,omitempty" name:"ExcludeNamespace"`

	// 需要排除的pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeLabels []*string `json:"ExcludeLabels,omitempty" name:"ExcludeLabels"`
}

type ContainerWorkLoadInfo struct {

	// 工作负载的类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 工作负载的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 容器名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ContentInfo struct {

	// 内容格式，支持json、csv
	Format *string `json:"Format,omitempty" name:"Format"`

	// csv格式内容描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Csv *CsvInfo `json:"Csv,omitempty" name:"Csv"`

	// json格式内容描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *JsonInfo `json:"Json,omitempty" name:"Json"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 通知类型。可选值：
	// <li> Trigger - 告警触发
	// <li> Recovery - 告警恢复
	// <li> All - 告警触发和告警恢复
	Type *string `json:"Type,omitempty" name:"Type"`

	// 通知接收对象。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// 接口回调信息（包括企业微信）。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警模板ID
		AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 触发条件。
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知模板列表。
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// 是否开启告警策略。默认值为true
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 用户自定义告警内容
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// 用户自定义回调
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`

	// 多维分析
	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AlarmTargets")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "Status")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略ID。
		AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志主题id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 类型：container_stdout、container_file、host_file
	Type *string `json:"Type,omitempty" name:"Type"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 采集配置标
	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`

	// 日志集id
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集name
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 节点文件配置信息
	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`

	// 容器文件路径信息
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`

	// 容器标准输出信息
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`

	// 日志格式化方式
	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// 用户自定义采集规则，Json格式序列化的字符串
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`

	// 绑定的机器组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 绑定的机器组id列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *CreateConfigExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TopicId")
	delete(f, "Type")
	delete(f, "LogType")
	delete(f, "ConfigFlag")
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	delete(f, "HostFile")
	delete(f, "ContainerFile")
	delete(f, "ContainerStdout")
	delete(f, "LogFormat")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	delete(f, "GroupId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigExtraResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采集配置扩展信息ID
		ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest

	// 采集配置名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 采集配置所属日志主题ID即TopicId
	Output *string `json:"Output,omitempty" name:"Output"`

	// 日志采集路径,包含文件名
	Path *string `json:"Path,omitempty" name:"Path"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// 用户自定义采集规则，Json格式序列化的字符串
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Output")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采集配置ID
		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 是否投递日志的元数据信息，默认为 true
	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
}

func (r *CreateConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "NeedContent")
	delete(f, "Content")
	delete(f, "Ckafka")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataTransformRequest struct {
	*tchttp.BaseRequest

	// 函数类型. DSL:1 SQL:2
	FuncType *int64 `json:"FuncType,omitempty" name:"FuncType"`

	// 源日志主题
	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`

	// 加工任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 加工逻辑函数
	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`

	// 加工任务目的topic_id以及别名
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`

	// 任务类型.  以SrcTopicId为数据源建立预览任务:1，以PreviewLogStatistics为数据源建立预览任务:2  真实任务:3
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务启动状态.   默认为1，正常开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`

	// 测试数据
	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitempty" name:"PreviewLogStatistics"`
}

func (r *CreateDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FuncType")
	delete(f, "SrcTopicId")
	delete(f, "Name")
	delete(f, "EtlContent")
	delete(f, "DstResources")
	delete(f, "TaskType")
	delete(f, "EnableFlag")
	delete(f, "PreviewLogStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志导出数量,  最大值5000万
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 日志导出检索语句，不支持<a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出起始时间，毫秒时间戳
	From *int64 `json:"From,omitempty" name:"From"`

	// 日志导出结束时间，毫秒时间戳
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitempty" name:"Format"`
}

func (r *CreateExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Count")
	delete(f, "Query")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Order")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出ID。
		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`

	// 是否生效，默认为true
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Rule")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集名字，不能重名
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 标签描述列表。最大支持10个标签键值对，并且不能有重复的键值对
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志集ID
		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组名字，不能重复
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建机器组类型，Type为ip，Values中为Ip字符串列表创建机器组，Type为label， Values中为标签字符串列表创建机器组
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器组ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperRequest struct {
	*tchttp.BaseRequest

	// 创建的投递规则所属的日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 创建的投递规则投递的bucket
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 创建的投递规则投递目录的前缀
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// 投递规则的名字
	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒，默认300，范围 300-900
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB，默认256，范围 100-256
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
}

func (r *CreateShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 投递规则ID
		ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志主题分区个数。默认创建1个，最大支持创建10个分区。
	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启自动分裂，默认值为true
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// 开启自动分裂后，每个主题能够允许的最大分区数，默认值为50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// 日志主题的存储类型，可选值 hot（实时存储），cold（低频存储）；默认为hot。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 生命周期，单位天，可取值范围1~3600。取值为3640时代表永久保存
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "Tags")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "StorageType")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsvInfo struct {

	// csv首行是否打印key
	PrintKey *bool `json:"PrintKey,omitempty" name:"PrintKey"`

	// 每列key的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*string `json:"Keys,omitempty" name:"Keys"`

	// 各字段间的分隔符
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`

	// 若字段内容中包含分隔符，则使用该转义符包裹改字段，只能填写单引号、双引号、空字符串
	EscapeChar *string `json:"EscapeChar,omitempty" name:"EscapeChar"`

	// 对于上面指定的不存在字段使用该内容填充
	NonExistingField *string `json:"NonExistingField,omitempty" name:"NonExistingField"`
}

type DataTransformResouceInfo struct {

	// 目标主题id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type DataTransformTaskInfo struct {

	// 数据加工任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据加工任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务启用状态，默认为1，正常开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`

	// 加工任务类型，1： DSL， 2：SQL
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 源日志主题
	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`

	// 当前加工任务状态（1准备中/2运行中/3停止中/4已停止）
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 加工任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 最后启用时间，如果需要重建集群，修改该时间
	LastEnableTime *string `json:"LastEnableTime,omitempty" name:"LastEnableTime"`

	// 日志主题名称
	SrcTopicName *string `json:"SrcTopicName,omitempty" name:"SrcTopicName"`

	// 日志集id
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 加工任务目的topic_id以及别名
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`

	// 加工逻辑函数
	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集规则扩展配置ID
	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
}

func (r *DeleteConfigExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigExtraId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigExtraResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFromMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 采集配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigFromMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFromMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFromMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigFromMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest

	// 采集规则配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataTransformRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *DeleteLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
}

func (r *DeleteShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// <li> name
	// 按照【通知渠道组名称】进行过滤。
	// 类型：String
	// 必选：否
	// <li> alarmNoticeId
	// 按照【通知渠道组ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li> uid
	// 按照【接收用户ID】进行过滤。
	// 类型：String
	// 必选：否
	// <li> groupId
	// 按照【接收用户组ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警通知模板列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitempty" name:"AlarmNotices"`

		// 符合条件的告警通知模板总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	// 
	// 按照【告警策略名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> alarmId
	// 
	// 按照【告警策略ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> topicId
	// 
	// 按照【监控对象的日志主题ID】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> enable
	// 
	// 按照【启用状态】进行过滤。
	// 
	// 类型：String
	// 
	// 备注：enable参数值范围: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False。 其它值将返回参数错误信息.
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略列表。
		Alarms []*AlarmInfo `json:"Alarms,omitempty" name:"Alarms"`

		// 符合查询条件的告警策略数目。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigExtrasRequest struct {
	*tchttp.BaseRequest

	// 支持的key： topicId,name, configExtraId, machineGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigExtrasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigExtrasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigExtrasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigExtrasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采集配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Configs []*ConfigExtraInfo `json:"Configs,omitempty" name:"Configs"`

		// 过滤到的总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigExtrasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigExtrasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// 采集配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采集规则配置绑定的机器组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest

	// <br><li> configName
	// 
	// 按照【采集配置名称】进行模糊匹配过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> configId
	// 
	// 按照【采集配置ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> topicId
	// 
	// 按照【日志主题】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采集配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`

		// 过滤到的总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 投递任务是否生效
		Effective *bool `json:"Effective,omitempty" name:"Effective"`

		// 是否投递日志的元数据信息
		NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`

		// 如果需要投递元数据信息，元数据信息的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Content *ConsumerContent `json:"Content,omitempty" name:"Content"`

		// CKafka的描述
		Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformInfoRequest struct {
	*tchttp.BaseRequest

	// <br><li> taskName
	// 
	// 按照【加工任务名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> taskId
	// 
	// 按照【加工任务id】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 默认值为2.   1: 获取单个任务的详细信息 2：获取任务列表
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Type为1， 此参数必填
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTransformInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataTransformInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据加工任务列表信息
		DataTransformTaskInfos []*DataTransformTaskInfo `json:"DataTransformTaskInfos,omitempty" name:"DataTransformTaskInfos"`

		// 任务总次数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTransformInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出列表
		Exports []*ExportInfo `json:"Exports,omitempty" name:"Exports"`

		// 总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 是否生效
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`

		// 索引修改时间，初始值为索引创建时间。
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志时间,  格式: YYYY-mm-dd HH:MM:SS.FFF
	BTime *string `json:"BTime,omitempty" name:"BTime"`

	// 日志包序号
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号
	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 上文日志条数,  默认值10
	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`

	// 下文日志条数,  默认值10
	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "BTime")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "PrevLogs")
	delete(f, "NextLogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志上下文信息集合
		LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitempty" name:"LogContextInfos"`

		// 上文日志是否已经返回
		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`

		// 下文日志是否已经返回
		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 时间间隔: 单位ms
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeLogHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogHistogramResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计周期： 单位ms
		Interval *int64 `json:"Interval,omitempty" name:"Interval"`

		// 命中关键字的日志总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 周期内统计结果详情
		HistogramInfos []*HistogramInfo `json:"HistogramInfos,omitempty" name:"HistogramInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsRequest struct {
	*tchttp.BaseRequest

	// <br><li> logsetName
	// 
	// 按照【日志集名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> logsetId
	// 
	// 按照【日志集ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tagKey
	// 
	// 按照【标签键】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tag:tagKey
	// 
	// 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// 类型：String
	// 
	// 必选：否
	// 
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLogsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志集列表
		Logsets []*LogsetInfo `json:"Logsets,omitempty" name:"Logsets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsRequest struct {
	*tchttp.BaseRequest

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachineGroupConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 采集规则配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// <br><li> machineGroupName
	// 
	// 按照【机器组名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> machineGroupId
	// 
	// 按照【机器组ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tagKey
	// 
	// 按照【标签键】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tag:tagKey
	// 
	// 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// 类型：String
	// 
	// 必选：否
	// 
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器组信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`

		// 分页的总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 查询的机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器状态信息组
		Machines []*MachineInfo `json:"Machines,omitempty" name:"Machines"`

		// 机器组是否开启自动升级功能
		AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

		// 机器组自动升级功能预设开始时间
		UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

		// 机器组自动升级功能预设结束时间
		UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

		// 当前用户可用最新的Loglistener版本
		LatestAgentVersion *string `json:"LatestAgentVersion,omitempty" name:"LatestAgentVersion"`

		// 是否开启服务日志
		ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分区列表
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// 查询的开始时间戳，支持最近3天的查询， 毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的结束时间戳， 毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeShipperTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShipperTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 投递任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Tasks []*ShipperTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShipperTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShippersRequest struct {
	*tchttp.BaseRequest

	// <br><li> shipperName
	// 
	// 按照【投递规则名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> shipperId
	// 
	// 按照【投递规则ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> topicId
	// 
	// 按照【日志主题】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeShippersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShippersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShippersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 投递规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Shippers []*ShipperInfo `json:"Shippers,omitempty" name:"Shippers"`

		// 本次查询获取到的总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShippersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// <br><li> topicName按照【日志主题名称】进行过滤。类型：String必选：否<br><li> logsetName按照【日志集名称】进行过滤。类型：String必选：否<br><li> topicId按照【日志主题ID】进行过滤。类型：String必选：否<br><li> logsetId按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。类型：String必选：否<br><li> tagKey按照【标签键】进行过滤。类型：String必选：否<br><li> tag:tagKey按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换，例如tag:exampleKey。类型：String必选：否<br><li> storageType按照【日志主题的存储类型】进行过滤。可选值 hot（实时存储），cold（低频存储）类型：String必选：否每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题列表
		Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`

		// 总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExcludePathInfo struct {

	// 类型，选填File或Path
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type对应的具体内容
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ExportInfo struct {

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志导出任务ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// 日志导出查询语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 日志导出时间排序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 日志导出格式
	Format *string `json:"Format,omitempty" name:"Format"`

	// 日志导出数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 日志导出起始时间
	From *int64 `json:"From,omitempty" name:"From"`

	// 日志导出结束时间
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出路径
	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`

	// 日志导出创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ExtractRuleInfo struct {

	// 时间字段的key名字，time_key和time_format必须成对出现
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// 时间字段的格式，参考c语言的strftime函数对于时间的格式说明输出参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// 分隔符类型日志的分隔符，只有log_type为delimiter_log时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`

	// 整条日志匹配规则，只有log_type为fullregex_log时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`

	// 行首匹配规则，只有log_type为multiline_log或fullregex_log时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginRegex *string `json:"BeginRegex,omitempty" name:"BeginRegex"`

	// 取的每个字段的key名字，为空的key代表丢弃这个字段，只有log_type为delimiter_log时有效，json_log的日志使用json本身的key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*string `json:"Keys,omitempty" name:"Keys"`

	// 需要过滤日志的key，及其对应的regex
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterKeyRegex []*KeyRegexInfo `json:"FilterKeyRegex,omitempty" name:"FilterKeyRegex"`

	// 解析失败日志是否上传，true表示上传，false表示不上传
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnMatchUpLoadSwitch *bool `json:"UnMatchUpLoadSwitch,omitempty" name:"UnMatchUpLoadSwitch"`

	// 失败日志的key
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnMatchLogKey *string `json:"UnMatchLogKey,omitempty" name:"UnMatchLogKey"`

	// 增量采集模式下的回溯数据量，默认-1（全量采集）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backtracking *int64 `json:"Backtracking,omitempty" name:"Backtracking"`
}

type Filter struct {

	// 需要过滤的字段。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 需要过滤的值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FilterRuleInfo struct {

	// 过滤规则Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 过滤规则
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 过滤规则Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type FullTextInfo struct {

	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 全文索引的分词符，字符串中每个字符代表一个分词符
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为1024
	Query *string `json:"Query,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 为true代表使用新检索,响应参数AnalysisRecords和Columns有效， 为false时代表使用老检索方式, AnalysisResults和ColNames有效
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加载后续内容的Context
		Context *string `json:"Context,omitempty" name:"Context"`

		// 日志查询结果是否全部返回
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 返回的是否为分析结果
		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

		// 如果Analysis为True，则返回分析结果的列名，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

		// 日志查询结果；当Analysis为True时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// 日志分析结果；当Analysis为False时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

		// 新的日志分析结果; UseNewAnalysis为true有效
	// 注意：此字段可能返回 null，表示取不到有效值。
		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`

		// 日志分析的列属性; UseNewAnalysis为true有效
	// 注意：此字段可能返回 null，表示取不到有效值。
		Columns []*Column `json:"Columns,omitempty" name:"Columns"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistogramInfo struct {

	// 统计周期内的日志条数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 按 period 取整后的 unix timestamp： 单位毫秒
	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
}

type HostFileInfo struct {

	// 日志文件夹
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// 日志文件名
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// metadata信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLabels []*string `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

type JsonInfo struct {

	// 启用标志
	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`

	// 元数据信息列表, 可选值为 __SOURCE__、__FILENAME__、__TIMESTAMP__。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`
}

type KeyRegexInfo struct {

	// 需要过滤日志的key
	Key *string `json:"Key,omitempty" name:"Key"`

	// key对应的过滤规则regex
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type KeyValueInfo struct {

	// 需要配置键值或者元字段索引的字段，元字段Key无需额外添加`__TAG__.`前缀，与上传日志时对应的字段Key一致即可，腾讯云控制台展示时将自动添加`__TAG__.`前缀
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段的索引描述信息
	Value *ValueInfo `json:"Value,omitempty" name:"Value"`
}

type LogContextInfo struct {

	// 日志来源设备
	Source *string `json:"Source,omitempty" name:"Source"`

	// 采集路径
	Filename *string `json:"Filename,omitempty" name:"Filename"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 日志包序号
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号
	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 日志时间戳
	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
}

type LogInfo struct {

	// 日志时间，单位ms
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志来源IP
	Source *string `json:"Source,omitempty" name:"Source"`

	// 日志文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志上报请求包的ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 请求包内日志的ID
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 日志内容的Json序列化字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
}

type LogItem struct {

	// 日志Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 日志Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogItems struct {

	// 分析结果返回的KV数据对
	Data []*LogItem `json:"Data,omitempty" name:"Data"`
}

type LogsetInfo struct {

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集名称
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 日志集绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 日志集下日志主题的数目
	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`

	// 若AssumerUin非空，则表示创建该日志集的服务方角色
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type MachineGroupInfo struct {

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器组类型
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 机器组绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

type MachineGroupTypeInfo struct {

	// 机器组类型，ip表示该机器组Values中存的是采集机器的IP地址，label表示该机器组Values中存储的是机器的标签
	Type *string `json:"Type,omitempty" name:"Type"`

	// 机器描述列表
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type MachineInfo struct {

	// 机器的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 机器状态，0:异常，1:正常
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 机器离线时间，空为正常，异常返回具体时间
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 机器是否开启自动升级。0:关闭，1:开启
	AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 机器当前版本号。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 机器升级功能状态。
	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 机器升级结果标识。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 机器升级结果信息。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 合并的PartitionId
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合并结果集
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组ID。
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

	// 通知渠道组名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 通知类型。可选值：
	// <li> Trigger - 告警触发
	// <li> Recovery - 告警恢复
	// <li> All - 告警触发和告警恢复
	Type *string `json:"Type,omitempty" name:"Type"`

	// 通知接收对象。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// 接口回调信息（包括企业微信）。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 触发条件。
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知模板列表。
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// 是否开启告警策略。
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 用户自定义告警内容
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// 用户自定义回调
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`

	// 多维分析
	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	delete(f, "Name")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "AlarmTargets")
	delete(f, "Status")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集配置扩展信息id
	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`

	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志主题id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 节点文件配置信息
	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`

	// 容器文件路径信息
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`

	// 容器标准输出信息
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 日志格式化方式
	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// 用户自定义采集规则，Json格式序列化的字符串
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`

	// 类型：container_stdout、container_file、host_file
	Type *string `json:"Type,omitempty" name:"Type"`

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 自建采集配置标
	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集name
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *ModifyConfigExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigExtraId")
	delete(f, "Name")
	delete(f, "TopicId")
	delete(f, "HostFile")
	delete(f, "ContainerFile")
	delete(f, "ContainerStdout")
	delete(f, "LogType")
	delete(f, "LogFormat")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	delete(f, "Type")
	delete(f, "GroupId")
	delete(f, "ConfigFlag")
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigExtraResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigRequest struct {
	*tchttp.BaseRequest

	// 采集规则配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 采集规则配置名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志采集路径，包含文件名
	Path *string `json:"Path,omitempty" name:"Path"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`

	// 采集配置关联的日志主题（TopicId）
	Output *string `json:"Output,omitempty" name:"Output"`

	// 用户自定义解析字符串，Json格式序列化的字符串
	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
}

func (r *ModifyConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Name")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "Output")
	delete(f, "UserDefineRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 投递任务是否生效，默认不生效
	Effective *bool `json:"Effective,omitempty" name:"Effective"`

	// 是否投递日志的元数据信息，默认为 false
	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
}

func (r *ModifyConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Effective")
	delete(f, "NeedContent")
	delete(f, "Content")
	delete(f, "Ckafka")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataTransformRequest struct {
	*tchttp.BaseRequest

	// 加工任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 加工任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 加工逻辑函数
	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`

	// 任务启动状态. 默认为1，正常开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`

	// 加工任务目的topic_id以及别名
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
}

func (r *ModifyDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "EtlContent")
	delete(f, "EnableFlag")
	delete(f, "DstResources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 默认不生效
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集名称
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志集的绑定的标签键值对。最大支持10个标签键值对，同一个资源只能同时绑定一个标签键。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器组类型
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// 投递规则投递的新的bucket
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 投递规则投递的新的目录前缀
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// 投递规则的开关状态
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 投递规则的名字
	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒，默认300，范围 300-900
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB，默认256，范围 100-256
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
}

func (r *ModifyShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "Status")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 该日志主题是否开始采集
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 是否开启自动分裂
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// 若开启最大分裂，该主题能够能够允许的最大分区数
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// 生命周期，单位天，可取值范围1~3600。取值为3640时代表永久保存
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TopicName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTime struct {

	// 可选值：
	// <br><li> Period - 周期执行
	// <br><li> Fixed - 定期执行
	Type *string `json:"Type,omitempty" name:"Type"`

	// 执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。
	Time *int64 `json:"Time,omitempty" name:"Time"`
}

type NoticeReceiver struct {

	// 接受者类型。可选值：
	// <br><li> Uin - 用户ID
	// <br><li> Group - 用户组ID
	// 暂不支持其余接收者类型。
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 接收者。
	ReceiverIds []*int64 `json:"ReceiverIds,omitempty" name:"ReceiverIds"`

	// 通知接收渠道。
	// <br><li> Email - 邮件
	// <br><li> Sms - 短信
	// <br><li> WeChat - 微信
	// <br><li> Phone - 电话
	ReceiverChannels []*string `json:"ReceiverChannels,omitempty" name:"ReceiverChannels"`

	// 允许接收信息的开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 允许接收信息的结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 位序
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type OpenKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS控制台创建的TopicId
	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *OpenKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 待消费TopicId
		TopicID *string `json:"TopicID,omitempty" name:"TopicID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PartitionInfo struct {

	// 分区ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`

	// 分区的状态（readwrite或者是readonly）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分区哈希键起始key
	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitempty" name:"InclusiveBeginKey"`

	// 分区哈希键结束key
	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitempty" name:"ExclusiveEndKey"`

	// 分区创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 只读分区数据停止写入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastWriteTime *string `json:"LastWriteTime,omitempty" name:"LastWriteTime"`
}

type PreviewLogStatistic struct {

	// 日志内容
	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`

	// 行号
	LineNum *int64 `json:"LineNum,omitempty" name:"LineNum"`

	// 目标日志主题
	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`

	// 失败错误码， 空字符串""表示正常
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 日志时间戳
	Time *string `json:"Time,omitempty" name:"Time"`

	// 目标topic-name
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstTopicName *string `json:"DstTopicName,omitempty" name:"DstTopicName"`
}

type RetryShipperTaskRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// 投递任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RetryShipperTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryShipperTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RetryShipperTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryShipperTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// 全文索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`

	// 键值索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`

	// 元字段索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *RuleTagInfo `json:"Tag,omitempty" name:"Tag"`
}

type RuleKeyValueInfo struct {

	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 需要建立索引的键值对信息；最大只能配置100个键值对
	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type RuleTagInfo struct {

	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 元字段索引配置中的字段信息
	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 要检索分析的日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要检索分析的日志的起始时间，Unix时间戳（毫秒）
	From *int64 `json:"From,omitempty" name:"From"`

	// 要检索分析的日志的结束时间，Unix时间戳（毫秒）
	To *int64 `json:"To,omitempty" name:"To"`

	// 检索分析语句，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a> | <a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>构成，无需对日志进行统计分析时，可省略其中的管道符<code> | </code>及SQL语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 表示单次查询返回的原始日志条数，最大值为1000，获取后续日志需使用Context参数
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果条数指定方式参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitempty" name:"Context"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 为true代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效
	// 为false时代表使用老的检索结果返回方式, 输出AnalysisResults和ColNames有效
	// 两种返回方式在编码格式上有少量区别，建议使用true
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时
		Context *string `json:"Context,omitempty" name:"Context"`

		// 符合检索条件的日志是否已全部返回，如未全部返回可使用Context参数获取后续更多日志
	// 注意：仅当检索分析语句(Query)不包含SQL时有效
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 返回的是否为统计分析（即SQL）结果
		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

		// 匹配检索条件的原始日志
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// 日志统计分析结果的列名
	// 当UseNewAnalysis为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

		// 日志统计分析结果
	// 当UseNewAnalysis为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

		// 日志统计分析结果
	// 当UseNewAnalysis为true时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`

		// 日志统计分析结果的列属性
	// 当UseNewAnalysis为true时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
		Columns []*Column `json:"Columns,omitempty" name:"Columns"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShipperInfo struct {

	// 投递规则ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 投递的bucket地址
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 投递的前缀目录
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// 投递规则的名字
	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 是否生效
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 投递日志的过滤规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *ContentInfo `json:"Content,omitempty" name:"Content"`

	// 投递日志的创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ShipperTaskInfo struct {

	// 投递任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 投递信息ID
	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 本批投递的日志的开始时间戳，毫秒
	RangeStart *int64 `json:"RangeStart,omitempty" name:"RangeStart"`

	// 本批投递的日志的结束时间戳， 毫秒
	RangeEnd *int64 `json:"RangeEnd,omitempty" name:"RangeEnd"`

	// 本次投递任务的开始时间戳， 毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 本次投递任务的结束时间戳， 毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 本次投递的结果，"success","running","failed"
	Status *string `json:"Status,omitempty" name:"Status"`

	// 结果的详细信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 待分裂分区ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`

	// 分区切分的哈希key的位置，只在Number=2时有意义
	SplitKey *string `json:"SplitKey,omitempty" name:"SplitKey"`

	// 分区分裂个数(可选)，默认等于2
	Number *int64 `json:"Number,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	delete(f, "SplitKey")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SplitPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分裂结果集
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TopicInfo struct {

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区个数
	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// 是否开启索引
	Index *bool `json:"Index,omitempty" name:"Index"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 日主主题是否开启采集
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 日志主题绑定的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 该主题是否开启自动分裂
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// 若开启自动分裂的话，该主题能够允许的最大分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// 日主题的存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 生命周期，单位天，可取值范围1~3600。取值为3640时代表永久保存
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type UploadLogRequest struct {
	*tchttp.BaseRequest

	// 主题id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 根据 hashkey 写入相应范围的主题分区
	HashKey *string `json:"HashKey,omitempty" name:"HashKey"`

	// 压缩方法
	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
}

func (r *UploadLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "HashKey")
	delete(f, "CompressType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValueInfo struct {

	// 字段类型，目前支持的类型有：long、text、double
	Type *string `json:"Type,omitempty" name:"Type"`

	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// 字段是否开启分析功能
	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type WebCallback struct {

	// 回调地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 回调的类型。可选值：
	// <li> WeCom
	// <li> Http
	CallbackType *string `json:"CallbackType,omitempty" name:"CallbackType"`

	// 回调方法。可选值：
	// <li> POST
	// <li> PUT
	// 默认值为POST。CallbackType为Http时为必选。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求头。
	// 注意：该参数已废弃，请在<a href="https://cloud.tencent.com/document/product/614/56466">创建告警策略</a>接口CallBack参数中指定请求头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*string `json:"Headers,omitempty" name:"Headers"`

	// 请求内容。
	// 注意：该参数已废弃，请在<a href="https://cloud.tencent.com/document/product/614/56466">创建告警策略</a>接口CallBack参数中指定请求内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Body *string `json:"Body,omitempty" name:"Body"`

	// 序号
	Index *int64 `json:"Index,omitempty" name:"Index"`
}
