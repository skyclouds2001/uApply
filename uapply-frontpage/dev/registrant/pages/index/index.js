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
		nice: false,
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
						// todo: request rewrite
						app.request({
							url: {
								key: 'user_login',
								arg: code.code
							},
							success(res){
								wx.hideLoading()
								Object.assign(app.user, res.data);
								thisPage.setData({
									app: getApp()
								})
								wx.setStorageSync('lst', Date.now());
								wx.setStorageSync('user', app.user);
							}
						})
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
		if(this.data.nice){
			cd_module(() => {
				wx.navigateTo({
					url: '/subpages/my_resume/my_resume',
				})
			})
		}else{
			wx.showToast({
				title: '暂未开放',
				icon: "none"
			})
		}
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
	onLoad(){
		const thisPage = this;
		app.request({
			url: {
				key: 'nice'
			},
			success(res){
				if (res.data.code >= app.code) {
					thisPage.setData({
						nice: true
					})
				}
			}
		})
	}
})
