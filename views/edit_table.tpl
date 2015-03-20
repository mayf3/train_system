{{template "head.tpl" .}}

<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<button class="btn btn-primary" type="button" onclick="window.location='/create_information'">Create New Team</button>
			<table class="table">
				<thead>
					<tr>
						<th>
							Team
						</th>
						{{ range $key, $val := .problem_list}}
						<th>
							{{ $val}}
						</th>
						{{end}}
					</tr>
				</thead>
				<tbody>
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
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
