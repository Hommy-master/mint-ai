package types

type TaskStatus string

const (
	TSPending   TaskStatus = "Pending"   // 等待执行
	TSRunning   TaskStatus = "Running"   // 正在执行
	TSCompleted TaskStatus = "Completed" // 执行完成
	TSFailed    TaskStatus = "Failed"    // 执行失败
)

type TaskType string

const (
	TTFrames        TaskType = "Frames"        // 帧任务
	TTConcat        TaskType = "Concat"        // 拼接任务
	TTExtraction    TaskType = "Extraction"    // 提取任务
	TTRemux         TaskType = "Remux"         // 重封装任务
	TTConvert       TaskType = "Convert"       // 转换任务
	TTTrim          TaskType = "Trim"          // 裁剪任务
	TTRescale       TaskType = "Rescale"       // 缩放任务
	TTCreate        TaskType = "Create"        // 创建任务
	TTAddImage      TaskType = "AddImage"      // 添加图片任务
	TTAddSubtitle   TaskType = "AddSubtitle"   // 添加字幕任务
	TTAddSubtitleEx TaskType = "AddSubtitleEx" // 添加字幕任务Ex
	TTAutoSubtitle  TaskType = "AutoSubtitle"  // 自动添加字幕任务
	TTAddText       TaskType = "AddText"       // 添加字幕任务
	TTAddTextEx     TaskType = "AddTextEx"     // 添加字幕任务Ex
	TTAutoText      TaskType = "AutoText"      // 自动添加字幕任务
	TTPIP           TaskType = "PIP"           // 画中画任务
)

type Task struct {
	TT  TaskType    // 任务类型
	TS  TaskStatus  // 任务状态
	R   interface{} // 任务结果
	Err error       // 任务返回错误信息
}
