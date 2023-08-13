// subpages/my_status/my_status.js
const app = getApp();

Page({
	data: {
		app,
		'user_register': "/user/register",
		department: [],
		pop: true
	},
	onShow(){
		const thispage = this;
		wx.request({
			url: app.apiurl('user_register'),
			header: {
				Authorization: app.user.accessToken
			},
			fail_back: true,
			success(res){
				if(res.statusCode !== 200){
					wx.showToast({
						title: `服务器错误 ${res.statusCode}`,
					})
				}else{
					if(res.data.code){
						wx.showToast({
						  	title: res.data.msg,
						})
					}else{
						thispage.department = res.data.rsp
					}
				}
			},
			fail(err){
				app._toast(err.errMsg, 'error')
			}
		})
	},
	on_pop(){
		this.setData({
			pop: !this.data.pop
		})
	}
})

