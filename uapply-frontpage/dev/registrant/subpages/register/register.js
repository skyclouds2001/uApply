const app = getApp()
Page({
	data: {
		app,
		title: "申报须知",
		alert_items: [
			"①各组织招新时段不同，仅处于其招新时段内的组织及部门可被查询到。若意向组织/部门未出现在列表，请耐心等待其招新启动/联系意向组织负责人确认",
			"②每个组织可申报至多一个志愿部门，已申报部门不可变更/取消",
			"③申报成功后请耐心等待面试通知，通知将通过短信/邮件发送，请确认当前简历中联系方式正确。",
			"④如因简历中联系方式有误等原因导致无法成功收到面试通知，请联系电话： 15209934027 解决"
		],
		current_organization: "",
		organization: [],
		pop: false
	},
	on_click_confirm_alert(){
		this.setData({
			pop: false
		})
	},
	on_click_alert(){
		this.setData({
			pop: true
		})
	},
	on_select(e){
		this.setData({
			current_organization: Number(e.detail.value)
		})
	},
	submit(e){
		if(!app.checkFormFull(e.detail.value, ['dep_id', 'org_id'])){
			wx.showToast({
				title: "请选选定目标组织及部门",
				icon: "none"
			})
			return
		}
		console.log(e.detail.value)
		app.request({
			url: {
				key: 'user_register'
			},
			method: "POST",
			data: {
				org_id: Number(e.detail.value.org_id),
				dep_id: Number(e.detail.value.dep_id),
			},
			accessToken: true,
			success(res){
				wx.showToast({
					title: res.data.msg,
					icon: "success"
				})
			},
		})
	},
	// 获取正在招新的组织和部门信息
	onLoad() {
		const thisPage = this
		wx.showLoading({
		  	title: '加载招新组织中',
			success(){
				app.request({
					url: {
						key: 'user_org-enrolling'
					},
					fail_back: true,
					success(res) {
						for (const [index, org] of res.data.organizations.entries()) {
							if (org['start_time'] < Date.now() < org['end_time']) {
								app.request({
									url: {
										key: 'user_org-deps',
										arg: org['org_id']
									},
									accessToken: true,
									fail_back: true,
									success(orgres) {
										thisPage.data.organization.push({
											id: org['org_id'],
											name: org['org_name'],
											department: orgres.data.deps
										})
										if(index === res.data.organizations.length - 1){
											wx.hideLoading({
												success: () => {
													thisPage.setData({
														organization: thisPage.data.organization
													})
													console.log(thisPage.data.organization)
												},
											})		
										}
									},
								})
							}
						}
					},
				})
			}
		})
	}
})