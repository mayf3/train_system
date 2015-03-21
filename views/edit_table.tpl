{{template "head.tpl" .}}

<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<button class="btn btn-primary" type="button" onclick="window.location='/create_information?table_id={{.table_id}}'">Create New Team</button>
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
						<th>
							operater
						</th>
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .total_status}}
						<tr class="success">
							<td>
								{{$val.Name}}
							</td>
							{{range $in_key, $in_val := $val.Status}}
								<td>
									{{$in_val}}
								</td>
							{{end}}
							<td>
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
