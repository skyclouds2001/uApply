Component({
	// 表单控件行为
	behaviors: ["wx://form-field"],
	relations: {
		"../selection/selection": {
			type: "child",
			linked(child) {
                // 如果设置了初始value，则尝试同步设置selected
				if(child.data.value === this.data.value){
					this.setData({
						selected: child.data.text
					})
				}
			}
		}
	},
	properties: {
		// 宽度
		width: {
			type: String,
			value: "auto"
		},
		// 背景颜色
		background: {
			type: String,
			value: "white"
		},
		// 颜色
		color: {
			type: String,
			value: "black"
        },
        // alue与所有selection不一致时（特别的，可能为没有设置value）的默认selection
		selected: {
            type: String,
            value: "NO_SELECTION"
        }
	},
	data: {
		// 是否展开显示其余选项
		stretch: false,
	},
	methods: {
		on_select(e){
			// 设置选中者与新值
			this.setData({
				value: e.detail.value,
				selected: e.detail.text
			})
			this.setData({
				stretch: !this.data.stretch
			})
		},
		on_stretch(){
			this.setData({
				stretch: !this.data.stretch
			})
		}
	},
})
