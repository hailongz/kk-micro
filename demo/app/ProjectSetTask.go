package app

/*B(Import)*/
	/*E(Import)*/

type /*B(Result)*/ ProjectSetTaskResult /*E(Result)*/ struct {
	/*B(Result.Base)*/
	/*E(Result.Base)*/

	/*B(Output)*/ /*E(Output)*/
	/*B(Output.project)*/
	Project *Project `json:"project,omitempty" title:"项目"`
	/*E(Output.project)*/
}

type /*B(Task)*/ ProjectSetTask /*E(Task)*/ struct {
	/*B(Task.Base)*/
	/*E(Task.Base)*/

	/*B(Input)*/ /*E(Input)*/
	/*B(Input.options)*/
	Options interface{} `json:"options" title:"其他属性"`
	/*E(Input.options)*/
	/*B(Input.status)*/
	Status interface{} `json:"status" title:"状态，200（上线），300（下线）"`
	/*E(Input.status)*/
	/*B(Input.endTime)*/
	EndTime interface{} `json:"endTime" title:"下线线时间（秒），0 为不自动下线"`
	/*E(Input.endTime)*/
	/*B(Input.inTime)*/
	InTime interface{} `json:"inTime" title:"上线时间（秒），0 为不自动上线"`
	/*E(Input.inTime)*/
	/*B(Input.maxCount)*/
	MaxCount interface{} `json:"maxCount" title:"最大数量 ，-1 为不限制"`
	/*E(Input.maxCount)*/
	/*B(Input.tags)*/
	Tags interface{} `json:"tags" title:"搜索关键字"`
	/*E(Input.tags)*/
	/*B(Input.title)*/
	Title interface{} `json:"title" title:"说明"`
	/*E(Input.title)*/
	/*B(Input.id)*/
	Id int64 `json:"id" title:"收款项目ID"`
	/*E(Input.id)*/

	/*B(Task.Result)*/
	Result ProjectSetTaskResult `json:"-"`
	/*E(Task.Result)*/
}

/*B(name)*/
func (T *ProjectSetTask) GetName() string {
	return "/project/set"
}

/*E(name)*/

/*B(title)*/
func (T *ProjectSetTask) GetTitle() string {
	return "修改项目"
}

/*E(title)*/

/*B(Task.GetResult)*/
func (T *ProjectSetTask) GetResult() interface{} {
	return &T.Result
}

/*E(Task.GetResult)*/
