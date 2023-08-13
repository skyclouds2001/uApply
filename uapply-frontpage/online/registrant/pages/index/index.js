// main.js
// 获取应用实例

const app = getApp()

// toast 请先登陆
function please_login_first(){
	wx.showToast({
		title: '未登录或获取ID失败',
		duration: 2000,
		icon: 'none'
	})
}

// 进入模块先判断是否登陆
function cd_module(fn){
	if(!app.user.uid){
		please_login_first()
	}else{
		fn()
	}
}

Page({
	data: {
		app,
	},
	on_click_get_id(){
		const thisPage = this;
		wx.getUserProfile({
			desc: '微信登录',
			success(res){
				app.user = res.userInfo
				thisPage.setData({
					app: getApp()
				})
				wx.login({
					success: function (code) {
						wx.showLoading({
							title: '登陆中',
						})
						wx.request({
							url: app.apiurl('user_login', code.code),
							success: function (res) {
								if(res.statusCode === 200){
									if (res.data.msg) {
										app.err_toast({
											errMsg: res.data.msg
										});
									}
									else {
										app._toast('登陆成功', 'success');
										Object.assign(app.user, res.data);
										thisPage.setData({
											app: getApp()
										})
										wx.setStorageSync('lst', Date.now());
										wx.setStorageSync('user', app.user);
									}
								}else{
									wx.showToast({
										title: "获取ID失败",
										icon: "error"
									})
								}
							},
							fail: function (err) {
								app.err_toast(err);
							}
						});
					},
					fail(err){
						wx.showToast({
							title: '获取ID失败',
							icon: "error"
						})
					}
				});
			},
		})
	},
	on_click_my_resume(){
		cd_module(() => wx.navigateTo({
			url: '/subpages/my_resume/my_resume',
		}))
	},
	on_click_register(){
		cd_module(() => {
			wx.navigateTo({
			  url: '/subpages/register/register',
			})
		})
	},
	on_click_my_status(){
		cd_module(() => wx.navigateTo({
			url: '/subpages/my_status/my_status',
		}))	
	},
	get_user_resume(){
		const thisPage = this;
		const user_resume = wx.getStorageSync(app.storage_list.user_resume)
		if (user_resume){
			app.user_resume = user_resume
		} else {
			if (!app.user){
				return
			}
			// 获取我的简历
			wx.request({
				url: app._translate(app.apis['user_check-resume']),
				headers: {
					Authentication: app.user.accessToken
				},
				success(res){
					this.user_resume = res.data.resume
				},
				fail(err){
					app.err_toast(err)
				}
			})
		}
	},
})
