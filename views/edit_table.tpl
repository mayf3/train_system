{{template "head.tpl" .}}

<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<button class="btn btn-primary" type="button" onclick="window.location='/create_information?table_id={{.table_id}}'">Create New Team</button>
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
						{{ range $key, $val := .problem_list}}
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
					{{range $key, $val := .total_status}}
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
							<td class="text-center">
								<button type="button" class="btn btn-sm btn-warning" onclick="window.location='/edit_information?table_id={{$.table_id}}&information_id={{$val.Id}}'">Edit</button>
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
