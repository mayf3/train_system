{{template "head.tpl" .}}

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
						{{if $.show}}
							<th>
                        	    Operation
							</th>
						{{end}}
					</tr>
				</thead>
				<tbody>
                    {{range $key, $val := .all_person}}
					<tr>
                        <td>
                            {{$val.Id}}
						</td>
					    <td>
                            {{$val.Name}}
						</td>
						<td>
                            {{$val.Grade}}
						</td>
						{{if $.show}}
							<td>
                        		<button type="button" class="btn btn-sm btn-info" onclick="window.location='/person/edit?person_id={{$val.Id}}'">Edit</button>
                            	<button type="button" class="btn btn-sm btn-warning" onclick="window.location='/person/delete?person_id={{$val.Id}}'">Delete</button>
							</td>
						{{end}}
					</tr>
                    {{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
