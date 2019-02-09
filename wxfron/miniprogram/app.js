
App({
  onLaunch: function() {
    //调用API从本地缓存中获取数据
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)
  },
  checkUserInfo: function(cb) {
    let that = this
    if (that.globalData.userInfo) {
      typeof cb == "function" && cb(that.globalData.userInfo, true);
    } else {
      console.log('no login')
      typeof cb == "function" && cb(that.globalData.userInfo, false);
    }
  },

  setLoginInfo:function(){
    var that = this
    var userInfo ={}
    wx.login({
      success: function (res) {
        if (res.code) {
          wx.getUserInfo({
            success: function (res) {
              userInfo.avatarUrl = res.userInfo.avatarUrl;
              userInfo.nickName = res.userInfo.nickName;
            }
          });
          var l = 'https://api.weixin.qq.com/sns/jscode2session?appid=' + that.globalData.appid + '&secret=' + that.globalData.appserct + '&js_code=' + res.code + '&grant_type=authorization_code';
          wx.request({
            url: l,
            data: {},
            method: 'GET', // OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, CONNECT  
            // header: {}, // 设置请求的 header  
            success: function (res) {
              //var obj = {};
              userInfo.openid = res.data.openid;
              wx.setStorageSync('userInfo', userInfo);//存储openid 
            }
          });
        } else {
          console.log('获取用户登录态失败！' + res.errMsg)
        }
      }
    });
  },
  globalData: {
    appid: "wx9b82b019e642a786",
    appserct: "fb97a7dc7f0dc3021f76d8a32663dbd9",
    openid: "",
    userInfo: null,
    //默认图片
    defaultImageUrl: 'https://www.wuchao.shop/static/self.jpg',
    imageUrl: 'http://image.bug2048.com/',
    imageStyle200To200: 'imageView2/1/w/200/h/200/q/100',
    imageStyle600To300: 'imageView2/1/w/600/h/300/q/75|watermark/2/text/QnVn55Sf5rS7MjA0OA==/font/5a6L5L2T/fontsize/280/fill/IzAwMDAwMA==/dissolve/100/gravity/SouthEast/dx/10/dy/10'
  }
})