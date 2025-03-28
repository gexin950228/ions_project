<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>知了内部办公网-后台系统管理</title>
	<meta name="renderer" content="webkit|ie-comp|ie-stand">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
    <meta http-equiv="Cache-Control" content="no-siteapp" />
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/static/css/font.css">
	<link rel="stylesheet" href="/static/css/xadmin.css">
    <link rel="icon" href="/static/images/zlkt.ico" type="image/x-icon"/>
    <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
    <script src="/static/lib/layui/layui.js" charset="utf-8"></script>
    <script type="text/javascript" src="/static/js/xadmin.js"></script>



    <link rel="stylesheet" href="/static/sweetalert/sweetalert.css">
    <script src="/static/sweetalert/ions_alert.js"></script>
    <script src="/static/sweetalert/sweetalert.min.js"></script>


</head>
<body class="login-bg">

    <div class="login">
        <div class="message">知了课堂内部办公网-登录管理</div>
        <div id="darkbannerwrap"></div>

        <form method="post" class="layui-form" >
            <input id="username" placeholder="用户名"  type="text" lay-verify="required" class="layui-input">
            <hr class="hr15">
            <input id="password" lay-verify="required" placeholder="密码"  type="password" class="layui-input">
            <hr class="hr15">
            <div>
                <input id="captcha" placeholder="验证码" type="text" lay-verify="required" class="layui-input" style="width: 180px;float: left">
                <img id="captcha_img" style="cursor:pointer;width: 140px;height: 50px;float: right"/>
                <input type="hidden" value="{{.captcha.Id}}" id="captcha_id">
            </div>
            <hr class="hr15">
            <input value="登录" lay-submit  lay-filter="login" style="width:100%;" type="button" id="btn">
            <hr class="hr20" >
        </form>
    </div>

    <script>
        var bs64 = {{.captcha.BS64}}
        document.getElementById("captcha_img").setAttribute("src",bs64);


        var img_bun = document.getElementById("captcha_img");
        var captcha_id = document.getElementById("captcha_id");
        img_bun.onclick = function (ev) {
            $.ajax({
                url:"/change_captcha",
                data:{},
                type:"GET",
                success:function (data) {
                    var code = data["Code"];
                    if(code != 200){  // 发生错误了
                        alert(data["msg"])
                    }
                    else {
                        var base64_value = data["BS64"];
                        var Id = data["Id"];
                        console.log(base64_value);
                        img_bun.setAttribute("src",base64_value);
                        captcha_id.setAttribute("value",Id);
                    }

                },
                fail:function (data) {

                }

            })
        };

        $(function  () {
            layui.use('form', function(){
              var form = layui.form;
              // layer.msg('玩命卖萌中', function(){
              //   //关闭后的操作
              //   });
              //监听提交
              form.on('submit(login)', function(data){
                  var username = document.getElementById("username").value;
                  var password = document.getElementById("password").value;
                  var captcha = document.getElementById("captcha").value;
                  var captcha_id = document.getElementById("captcha_id").value;

                  if(password.length < 6){
                      alert("密码长度不能少于6位");
                      return
                  }
                  $.ajax({
                      url:"/",
                      type:"POST",
                      data:{
                          "username":username,
                          "password":password,
                          "captcha":captcha,
                          "captcha_id":captcha_id
                      },
                      success:function (data) {
                          let code = data["code"];
                          console.log(code, data);
                          if(code === "200"){
                              // alert(data["msg"])
                              alert("haha")
                          }else {
                              alert("hehe")
                          }
                      },
                      fail:function (data) {
                        ions_alert.alertError(data["msg"])
                      }
                  })

              });
            });
        })


    </script>


    <!-- 底部结束 -->

</body>
</html>