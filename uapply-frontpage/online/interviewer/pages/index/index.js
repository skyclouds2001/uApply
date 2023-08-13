// main.js
// 获取应用实例

const app = getApp()
const delay = 1000

Page({
	data: {
		app,
        pop: false,
		organization: {},
		// 是否成功获取到组织和部门
	},
	move_module_in(){
        this.setData({
            st: {
                id_top: this.st.id_top.top(0).step().export(),
                md_org: this.st.md_org.left(0).step().export(),
                md_dep: this.st.md_dep.right(0).step().export(),
            }
        })
	},
	move_button_in(){
		this.setData({
			st: Object.assign(this.data.st, {
				btn_left: this.st.btn_left.right(0).step().export(),
				btn_right: this.st.btn_right.left(0).step().export()
			})
		})
	},
	onsubmit(e){
		if(!this.data.pop){
			this.get_registors(e.detail.value)
		}else{
			this.get_resumes(e.detail.value)
		}
	},
	get_registors(data){
		const thisPage = this;
		wx.navigateTo({
            url: "/subpages/registors/registors",
            success(target){
                target.eventChannel.emit('getData', {
                    depId: thisPage.data.organization.dep_id,
					dep_name: thisPage.data.organization.dep_name,
					org_name: thisPage.data.organization.org_name,
                })
                wx.setNavigationBarTitle({
                    title: '招新数据',
                })
            }
        })
	},
	get_resumes(data){
        const thispage = this;

		wx.navigateTo({
			url: "/subpages/resumes/resumes",
			success(target){
				thispage.setData({
					pop: false
				})
				target.eventChannel.emit('id', {
					uid: Number(data.uid),
					depId: Number(thispage.data.organization.dep_id)
				})
			}
		})
	},
	on_ent_pop(){
		this.setData({
			pop: true
		})
	},
	on_ext_pop(){
		this.setData({
			pop: false
		})
	},
	on_click_get_id() {
		const thisPage = this;
		wx.getUserProfile({
			desc: '微信登录',
			success(res) {
				app.user = res.userInfo
				thisPage.setData({
					app: getApp()
				})
				wx.login({
					success: function (code) {
						wx.showLoading({
							title: '登陆中',
						})
						app.request({
							url: {
								key: 'user_login',
								arg: String(code.code)
							},
							success(res){
								Object.assign(app.user, res.data);
								// 别删这行代码，否则你将为自己的自作聪明付出代价
								thisPage.setData({
									app: getApp()
								})
								// 开始获取面试官招新数据
								wx.showLoading({
									title: '加载招新组织中',
									success(){
										app.request({
											url : {
												key: 'inter_gerorgdep'
											},
											accessToken: true,
											success(res){
												wx.hideLoading({
													success(){
														let dep_id, org_id, dep_name, org_name;
														for(const dep in res.data.depId){
															dep_id = dep;
															dep_name = res.data.depId[dep]
														}
														for(const org in res.data.orgId){
															org_id = org;
															org_name = res.data.orgId[org]
														}
														thisPage.setData({
															organization: {
																org_id: org_id,
																org_name: org_name,
																dep_id: dep_id,
																dep_name: dep_name
															}
														})
														thisPage.move_module_in()
														thisPage.move_button_in()
													}
												})
											}
										})
									}
								})
							}
						})
					},
					fail(err) {
						wx.showToast({
							title: '获取ID失败',
							icon: "error"
						})
					}
				});
			},
		})
	},
	onLoad() {
        this.st = {
			id_top: wx.createAnimation({
				delay: 0,
                duration: delay,
                timingFunction: 'ease-out'
			}),
			md_org: wx.createAnimation({
				delay,
                duration: delay,
                timingFunction: 'ease-out'
			}),
			md_dep: wx.createAnimation({
				delay,
                duration: delay,
                timingFunction: "ease-out"
            }),
			// 招新数据按钮
			btn_left: wx.createAnimation({
				duration: delay,
				timingFunction: 'ease-out'
			}),
			// 获取简历按钮
			btn_right: wx.createAnimation({
				duration: delay,
				timingFunction: 'ease-out'
			})
		}
	},
})

