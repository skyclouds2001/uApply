// components/selection/selection.js
Component({
	relations: {
		"../select/select": {
			type: "parent"
		}	
	},
	properties: {
		// 值
		value: {
			type: String,
			value: ""
		},
		// 显示文本
		text: {
			type: String,
			value: ""
		}
	},
	data: {
	},
	methods: {
		on_tap(e){
			this.triggerEvent("select", {
				text: this.data.text,
				value: this.data.value
			}, {
				bubbles: true,
				composed: true,
				capturePhase: true
			})
		}
	}
})
