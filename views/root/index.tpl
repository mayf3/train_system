{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>SYSU Train System</title>
{{end}}

{{define "body"}}
<p></p>
<div class="container" role="main">
	<div class="row-fluid">
		<div class="span12">
			{{if compare $.permission "admin"}}
		    	<button class="btn btn-primary" type="button" onclick="window.location='{{ .url_create_table }}'">Create Table</button>
			{{end}}
            <p></p>
			<table class="table table-striped">
				<thead>
					<tr>
						<th>
							Date
						</th>
						<th>
							Title
						</th>
						<th>
                            Problem Source
						</th>
						<th>
                            Operation
						</th>
					</tr>
				</thead>
				<tbody>
                    {{range $key, $val := .show_all_table}}
					<tr>
                        <td>
                            {{$val.CreateTime}}
						</td>
						
					    <td>
                            {{$val.ContestName}}
						</td>
						<td>
                            {{$val.Source}}
						</td>
						<td>
                           <button type="button" class="btn btn-sm btn-info" onclick="window.location='/table/{{$val.Id}}/show'">Show</button>
                            <button type="button" class="btn btn-sm btn-warning" onclick="window.location='/table/{{$val.Id}}/edit'">Edit Team</button>
							{{if compare $.permission "admin"}}
                            	<button type="button" class="btn btn-sm btn-success" onclick="window.location='/table/{{$val.Id}}/setting'">Setting Team</button>
                            	<button type="button" class="btn btn-sm btn-warning" onclick="window.location='/table/{{$val.Id}}/del'">Delete</button>
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
