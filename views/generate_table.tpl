{{template "head.tpl" .}}

<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<h2 class="text-center">
				{{.title}}
			</h2>
			<h3>
				{{.source}}
			</h3>
			<h4 class="text-right">
				{{.date}}
			</h4>
			<ul>
				<li>
					新闻资讯
				</li>
				<li>
					体育竞技
				</li>
				<li>
					娱乐八卦
				</li>
				<li>
					前沿科技
				</li>
				<li>
					环球财经
				</li>
				<li>
					天气预报
				</li>
				<li>
					房产家居
				</li>
				<li>
					网络游戏
				</li>
			</ul>
			<table class="table table-bordered">
				<thead>
					<tr>
						<th>
							编号
						</th>
						<th>
							产品
						</th>
						<th>
							交付时间
						</th>
						<th>
							状态
						</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>
							1
						</td>
						<td>
							TB - Monthly
						</td>
						<td>
							01/04/2012
						</td>
						<td>
							Default
						</td>
					</tr>
					<tr class="success">
						<td>
							1
						</td>
						<td>
							TB - Monthly
						</td>
						<td>
							01/04/2012
						</td>
						<td>
							Approved
						</td>
					</tr>
					<tr class="error">
						<td>
							2
						</td>
						<td>
							TB - Monthly
						</td>
						<td>
							02/04/2012
						</td>
						<td>
							Declined
						</td>
					</tr>
					<tr class="warning">
						<td>
							3
						</td>
						<td>
							TB - Monthly
						</td>
						<td>
							03/04/2012
						</td>
						<td>
							Pending
						</td>
					</tr>
					<tr class="info">
						<td>
							4
						</td>
						<td>
							TB - Monthly
						</td>
						<td>
							04/04/2012
						</td>
						<td>
							Call in to confirm
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
