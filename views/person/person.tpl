{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Person - SYSU Train System</title>
{{end}}

{{define "body"}}
<script type="text/javascript">
$(document).ready(function(){
	$("#nav_index").attr("class", "");
	$("#nav_person").attr("class", "active");
});
</script>

<p></p>
<div class="container" role="main">
	<div class="row-fluid">
		<div class="span12">
			<form id="table" method="post" action="/person/add">
				<fieldset>
					<label>Name</label>
					<input name="name" type="text">
					<label>Grade</label>
					<input name="grade" type="text">
					<button type="submit" class="btn btn-success">Add New Person</button>
				</fieldset>
			</form>
            <p></p>
			<table class="table table-striped table-hover table-bordered">
				<thead>
					<tr>
						<th>
							#
						</th>
						<th>
							Name
						</th>
						<th>
							Grade
						</th>
						{{if compare $.permission "admin"}}
							<th>
                        	    Operation
							</th>
						{{end}}
					</tr>
				</thead>
				<tbody>
                    {{range $key, $val := .show_all_person}}
					<tr>
						<td class="text-center">
                            {{$val.Id}}
						</td>
						<td class="text-center">
                            {{$val.Name}}
						</td>
						<td class="text-center">
                            {{$val.Grade}}
						</td>
						{{if compare $.permission "admin"}}
							<td class="text-center">
								<div class="row">
									<div class="col-md-6">
										<form id="table" method="post" action="/person/{{$val.Id}}/edit">
                        					<button type="submit" class="btn btn-sm btn-info">Edit</button>
										</form>
									</div>
									<div class="col-md-6">
										<form id="table" method="post" action="/person/{{$val.Id}}/del">
                            				<button type="submit" class="btn btn-sm btn-warning">Del</button>
										</form>
									</div>
							</td>
						{{end}}
					</tr>
                    {{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>
{{end}}
