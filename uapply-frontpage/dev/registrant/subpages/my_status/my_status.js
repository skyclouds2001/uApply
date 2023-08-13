// subpages/my_status/my_status.js
const app = getApp();

Page({
	data: {
		app,
		department: [],
		pop: true
	},
	onShow(){
		const thispage = this;
		wx.showLoading({
			title: '获取数据中',
			success(){
				app.request({
					url: {
						key: 'user_register',
					},
					accessToken: true,
					success(res){
						wx.hideLoading({
							success: () => {
								thispage.setData({
									department: res.data.rsp,
									pop: res.data.rsp.length ? false : true
								})
								console.log(res)
							},
						})
					}
				})
			}
		})
	}
})

