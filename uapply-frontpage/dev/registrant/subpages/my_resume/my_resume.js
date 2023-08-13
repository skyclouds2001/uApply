// pages/registors.js


const app = getApp()
Page({
	data: {
		app: getApp(),
		alert: false,
		alert_items: [
			"①安全隐私说明：平台侧承诺您所填写的信息仅用于支持招新面试，不用于其他商业用途",
			"②内容填写完整提示：请尽可能将简历填写完整，更完善的简历会给面试官留下更好的印象哦",
			"③申报简历规则：每次申报时提交的简历以最新保存的简历为准",
			"④简历可重复编辑，请注意保存"
		],
		title: "简历填写须知",
		pop: false,
		resume: null,
		list: [
			{value: '男', text: '男'},
			{value: '女', text: '女'}
		],
		ok: false
	},
	on_click_alert(){
		this.setData({
			pop: true
		})
	},
	on_click_confirm_alert(){
		this.setData({
			pop: false,
			ok: true
		})
	},
	on_click_cancel_alert(){
		this.setData({
			pop: false
		})
		wx.navigateBack();
	},
	// 加载用户简历
	onLoad(){
		const thisPage = this;
		wx.enableAlertBeforeUnload({
			message: '简历尚未保存，若退出更新内容将丢失，是否确认退出简历页面填写页面',
		})
		app.request({
			url: {
				key: 'user_check-resume'
			},
			accessToken: true,
			success(res){
				thisPage.setData({
					resume: res.data.resume
				})
			}
		})
	},
	submit(e){
		const thisPage = this;
		if(!e.detail.value.read){
			wx.showToast({
				title: "请同意用户协议",
				icon: "error"
			})
			return
		}
		for (const i in e.detail.value){
			if(!e.detail.value[i]){
				wx.showToast({
					title: '请完善简历后保存',
					icon: "error"
				})
				return
			}
		}
		app.request({
			url: {
				key: 'user_resume',
			},
			accessToken: true,
			data: e.detail.value,
			fail_back: true,
			success() {
				wx.showToast({
					title: '保存简历成功',
					icon: "success"
				})
				wx.disableAlertBeforeUnload()
			}
		})
	}
})