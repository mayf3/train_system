{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Show Table- SYSU Train System</title>
{{end}}

{{define "body"}}
<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<h1 class="text-center">
				{{.show_current_table.ContestName}}
			</h1>
			
        	<hr></hr>

			<h4>
				{{.show_current_table.Source}}
			</h4>

			<h4 class="text-right">
				{{.show_current_table.CreateTime}}
			</h4>

        	<hr></hr>

			<ul>
				{{range $key, $val := .show_total_status}}
					<h4>
						<li>
							<strong>
								{{$val.Name}} :
							</strong>
							{{range $in_key, $in_val := .MemberName }}
								{{$in_val}}
							{{end}}
						</li>
					</h4>
				{{end}}
			</ul>
        	<hr></hr>
			<ul class="list-inline">
				<li class="text-muted">
					<h3>
						未AC
					</h3>
				</li>
				<li class="text-success">
					<h3>
						AC
					</h3>
				</li>
				<li class="text-danger">
					<h3>
						赛后AC
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
									{{$in_val.Context}}
								</td>
							{{end}}
						</tr>
					{{end}}
				</tbody>
			</table>

        	<hr></hr>

			<table class="table table-striped table-hover table-bordered">
				{{range $key, $val := .show_hust_table }}
					<tr>
						{{range $in_key, $in_val := $val}}
							<td>
								{{$in_val}}
							</td>
						{{end}}
					</tr>
				{{end}} 
			</table>
			{{if compare $.permission "admin"}}
				<form id="table" method="post" action="#">
					<input name="hust_table" type="input">
					<button type="submit" class="btn btn-success">Submit</button>
				</form>
			{{end}}
		</div>
	</div>
</div>
{{end}}
