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
                            {{$val.create_time}}
						</td>
						<td>
                            {{$val.contest_name}}
						</td>
						<td>
                            {{$val.source}}
						</td>
						<td>
                            <button type="button" class="btn btn-sm btn-success">Edit</button>
                            <button type="button" class="btn btn-sm btn-info">Show</button>
							Default
						</td>
					</tr>
                    {{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
