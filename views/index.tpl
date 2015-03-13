{{template "head.tpl" .}}

<p></p>
<div class="container" role="main">
	<div class="row-fluid">
		<div class="span12">
		    <button class="btn btn-primary" type="button" onclick="window.location='/create_table'">Create Table</button>
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
                    {{range $key, $val := .Map}}
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
                           <button type="button" class="btn btn-sm btn-info" onclick="window.location='/generate_table?table_id={{$val.Id}}'">Show</button>
                            <button type="button" class="btn btn-sm btn-success" onclick="window.location='/create_table?table_id={{$val.Id}}'">Edit_Table</button>
                            <button type="button" class="btn btn-sm btn-warning" onclick="window.location='/edit_table?table_id={{$val.Id}}'">Edit_Team</button>
                            <button type="button" class="btn btn-sm btn-warning" onclick="window.location='/action?action=DeleteTable&id={{$val.Id}}'">Delete</button>
						</td>
					</tr>
                    {{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
