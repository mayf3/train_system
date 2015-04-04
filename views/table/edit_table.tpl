{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Edit Table- SYSU Train System</title>
{{end}}

{{define "body"}}
<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<button class="btn btn-lg btn-primary" type="button" onclick="window.location='{{.url_create_information}}'">Create New Team</button>
			<ul class="list-inline">
				<li class="text-muted">
					<h3>
						<strong>
							未AC
						</strong>
					</h3>
				</li>
				<li class="text-success">
					<h3>
						<strong>
							AC
						</strong>
					</h3>
				</li>
				<li class="text-danger">
					<h3>
						<strong>
							赛后AC
						</strong>
					</h3>
				</li>
			</ul>
			<table class="table table-striped table-hover table-bordered">
				<thead>
					<tr>
						<th class="text-center">
							Team
						</th>
						{{ range $key, $val := .other_problem_list}}
						<th class="text-center">
							{{ $val}}
						</th>
						{{end}}
						<th class="text-center">
							operation
						</th>
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .show_total_status}}
						<tr>
							<td class="text-center">
								<strong>
								{{$val.Name}}
								</strong>
							</td>
							{{range $in_key, $in_val := $val.Status}}
								<td class="text-center {{$in_val.Colour}}">
						<strong>
									{{$in_val.Context}}
						</strong>
								</td>
							{{end}}
							<td class="text-center">
							<button type="button" class="btn btn-warning" onclick="window.location='{{$.url_current_table}}/information/{{$val.Id}}/edit'">Edit</button>
							{{if compare $.permission "admin"}}
								<button type="button" class="btn btn-danger" onclick="window.location='{{$.url_current_table}}/information/{{$val.Id}}/del'">Del</button>
							{{end}}
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>
{{end}}
