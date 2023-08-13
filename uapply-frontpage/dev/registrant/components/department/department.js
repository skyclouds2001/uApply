const status = ["新投递", "已经安排面试", "已面试", "通过", "淘汰"];

Component({
	properties: {
		background: {
			type: String,
			value: "white"
		},
		indexplus: {
			type: Number,
			value: 0
		},
		icon_bg:{
			type: String,
			value: "Black"
		},
		dep:{
			type: Object,
			value: {}
		}
	},
	data: {
		status: ""
	},
	methods: {
		translate_status(statu){
			return status[statu]
		}
	},
	options:{
		styleIsolation: 'apply-shared'
	},
	lifetimes:{
		attached(){
			this.setData({
				status : 
					this.data.dep.first ? 
					"一面 ► " + this.translate_status(this.data.dep.first) 
					: this.data.dep.second ? 
					"二面 ► " + this.translate_status(this.data.dep.second)
					: this.data.dep.final ? 
					"最终 ► " + this.translate_status(this.data.dep.final) 
					: "已投递"
			})
		}
	}
})
