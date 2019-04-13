package qyproxy

import (
	"fmt"
	"testing"
)

func TestProxyHostParse(t *testing.T) {

	data := `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>免费代理ip - 旗云代理 - 高品质Http代理ip供应平台/每天更新最新免费代理ip</title>
		<meta name="keywords" content="旗云代理,代理ip,高效ip提取,最新http代理,私密代理ip,爬虫代理ip,动态代理ip,长效代理ip,国内外代理ip,代理api接口">
    	<meta name="description" content="旗云代理是一个高质量http代理ip供应平台，拥有大量的ip资源其中包括免费代理、私密代理、开放代理、长效代理等多种类型的http和https代理ip；并且我们一直在探索更好的ip为用户提供旗舰级的代理服务，是您做爬虫与大数据事业的好帮手"/>
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
		<!-- Stylesheets -->
		<link rel="stylesheet" href="/css/font-awesome.min.css">
		<link rel="stylesheet" href="/css/main.min.css">
		<link rel="stylesheet" href="/css/loge.css">
	</head>
	<body class="footer-dark">
		<!-- Header -->
		<header id="header" class="header-dynamic header-shadow-scroll">
			<div class="container">
				<a class="logo" href="/">
					<img src="/img/logo.png" alt="">
				</a>
				<nav>
					<ul class="nav-primary">
						<li>
							<a href="/">首页</a>
						</li>
						
						<li>
							<a href="/free">免费代理</a>
						</li>
						
						<li>
							<a href="/buy">购买代理</a>
						</li>
						
						<li>
							<a href="/secret">私密代理</a>
						</li>
						
						<li>
							<a href="/putong">开放代理</a>
						</li>
						<li>
							<a>API文档</a>
							<ul style="min-width: 48px;">
								<li>
									<a href="/doc/putong/">开放代理</a>
								</li>
								<li>
									<a href="/doc/secret/">私密代理</a>
								</li>
							</ul>
						</li>

						<li>
							<a class="button button-secondary" href="/user/">
								<i class="fa fa-lock icon-left"></i>用户中心
							</a>
						</li>
					</ul>
					<ul class="nav-secondary">
						<li>
							<a target="_blank" href="http://wpa.qq.com/msgrd?v=3&uin=1138026702&site=qq&menu=yes"><i class="fa fa-comments icon-left"></i>咨询客服</a>
						</li>
					</ul>
				</nav>
			</div>
		</header>
		<section id="content">
			<section class="content-row content-row-gray">
			    <div align="center" style="padding:30px;">
			    <a class="button" style="background-color:#e6e9ee">国内代理</a>
			    <a class="button" href="?action=unchina">非国内代理</a>
				</div>
				<div class="container">
					<header class="content-header">
						<p>
							注：免费代理每小时少量更新，不代表本站收费代理质量。仅供开发者测试程序或参考使用！
						</p>
					</header>
				    <table class="table table-bordered table-striped">
          <thead>
              <tr>
                <th>IP</th>
                <th>PORT</th>
                <th>匿名度</th>
                <th>类型</th>
                <th>位置</th>
                <th>响应速度</th>
                <th>最后验证时间</th>
              </tr>
            </thead>
            <tbody>
                                                <tr>
                    <td data-title="IP">116.209.57.203</td>
                    <td data-title="PORT">9999</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTPS</td>
                    <td data-title="位置">中国 湖北 仙桃</td>
                    <td data-title="响应速度">0.586446秒</td>
                    <td data-title="最后验证时间">2019-04-13 09:51:27</td>
                </tr>
                                <tr>
                    <td data-title="IP">116.209.59.3</td>
                    <td data-title="PORT">9999</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTPS</td>
                    <td data-title="位置">中国 湖北 仙桃</td>
                    <td data-title="响应速度">0.59659秒</td>
                    <td data-title="最后验证时间">2019-04-13 09:59:16</td>
                </tr>
                                <tr>
                    <td data-title="IP">116.209.58.81</td>
                    <td data-title="PORT">9999</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTP</td>
                    <td data-title="位置">中国 湖北 仙桃</td>
                    <td data-title="响应速度">2.00143秒</td>
                    <td data-title="最后验证时间">2019-04-13 08:58:26</td>
                </tr>
                                <tr>
                    <td data-title="IP">122.117.190.229</td>
                    <td data-title="PORT">57876</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTP</td>
                    <td data-title="位置">中国 台湾 台南市</td>
                    <td data-title="响应速度">17.651543秒</td>
                    <td data-title="最后验证时间">2019-04-13 08:56:58</td>
                </tr>
                                <tr>
                    <td data-title="IP">121.127.234.141</td>
                    <td data-title="PORT">3128</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTPS</td>
                    <td data-title="位置">中国 香港</td>
                    <td data-title="响应速度">11.796481秒</td>
                    <td data-title="最后验证时间">2019-04-13 07:42:40</td>
                </tr>
                                <tr>
                    <td data-title="IP">118.144.149.206</td>
                    <td data-title="PORT">3128</td>
                    <td data-title="匿名度">透明</td>
                    <td data-title="类型">HTTP</td>
                    <td data-title="位置">中国 北京 北京</td>
                    <td data-title="响应速度">1.405653秒</td>
                    <td data-title="最后验证时间">2019-04-13 07:52:15</td>
                </tr>
                                <tr>
                    <td data-title="IP">116.209.57.102</td>
                    <td data-title="PORT">9999</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTPS</td>
                    <td data-title="位置">中国 湖北 仙桃</td>
                    <td data-title="响应速度">1.917602秒</td>
                    <td data-title="最后验证时间">2019-04-13 06:58:33</td>
                </tr>
                                <tr>
                    <td data-title="IP">116.209.57.207</td>
                    <td data-title="PORT">9999</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTPS</td>
                    <td data-title="位置">中国 湖北 仙桃</td>
                    <td data-title="响应速度">0.290611秒</td>
                    <td data-title="最后验证时间">2019-04-13 07:00:20</td>
                </tr>
                                <tr>
                    <td data-title="IP">39.137.77.67</td>
                    <td data-title="PORT">8080</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTP</td>
                    <td data-title="位置">中国 山东 烟台</td>
                    <td data-title="响应速度">0.10486秒</td>
                    <td data-title="最后验证时间">2019-04-13 05:48:12</td>
                </tr>
                                <tr>
                    <td data-title="IP">116.209.54.120</td>
                    <td data-title="PORT">9999</td>
                    <td data-title="匿名度">高匿</td>
                    <td data-title="类型">HTTP</td>
                    <td data-title="位置">中国 湖北 仙桃</td>
                    <td data-title="响应速度">2.190946秒</td>
                    <td data-title="最后验证时间">2019-04-13 06:01:55</td>
                </tr>
                                
            </tbody>
        </table>
        <br>
                    <nav aria-label="Page navigation" align="center">
                        <ul class="pagination">
                            <li><a href="?action=china&page=1" aria-label="Previous"><span aria-hidden="true">«</span></a></li>
                                                        <li><a href="?action=china&page=2">2</a></li>
                                                        <li><a href="?action=china&page=3">3</a></li>
                                                        <li><a href="?action=china&page=4">4</a></li>
                                                        <li><a href="?action=china&page=5">5</a></li>
                                                        <li><a href="?action=china&page=6">6</a></li>
                                                        <li><a href="?action=china&page=7">7</a></li>
                                                        <li><a href="?action=china&page=3" aria-label="Next"><span aria-hidden="true">»</span></a></li>
                        </ul>
                    </nav>
				</div>
			</section>
			<!-- Content Row -->
					</section>
		<!-- Footer -->
		<footer id="footer">
			<section class="footer-primary">
				<div class="container">
					<div class="column-row">
						<div class="column-33">
							<h5>
								关于我们
							</h5>
							<p>
								声明：本站资源仅限用来计算机技术学习及大数据抓取、爬虫研究等合法行为！利用本站资源从事任何违反法律法规的行为，由此引起的一切后果与本站无关。<br>
							</p>
						</div>
						<div class="column-66">
							<div class="column-row align-right-top">
								<div class="column-25">
									<h5>
										联系
									</h5>
									<ul class="list-style-icon">
										<li>
											<a target="_blank" href="http://wpa.qq.com/msgrd?v=3&uin=1138026702&site=qq&menu=yes"><i class="fa fa-qq"></i>QQ客服</a>
										</li>
									</ul>
								</div>
								<div class="column-25">
									<h5>
										产品
									</h5>
									<ul>
										<li>
											<a href="/secret">私密代理</a>
										</li>
										<li>
											<a href="/putong">开放代理</a>
										</li>
										<li>
											<a href="/#">长效代理</a>
										</li>
									</ul>
								</div>
							</div>
						</div>
					</div>
				</div>
			</section>
			<section class="footer-secondary">
				<div class="container">
					<p>
						Copyright 2019 &copy; qydaili.com 版权所有 粤ICP备18122265号-2 <script src="https://s19.cnzz.com/z_stat.php?id=1275082454&web_id=1275082454" language="JavaScript"></script>
					</p>
				</div>
			</section>
		</footer>
		<!-- Scripts -->
		<script>
            var contextPath = '';
        </script>
		<script src="/js/jquery.min.js"></script>
		<script src="/js/headroom.min.js"></script>
		<script src="/js/main.min.js"></script>
	</body>
</html>
`

	parse := ProxyHostParse([]byte(data), "")

	s := parse.NextRequest.Url
	fmt.Println(s)
}
