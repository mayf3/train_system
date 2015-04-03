{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Edit Information - SYSU Train System</title>
{{end}}

{{define "body"}}
<script type="text/javascript">
$(document).ready(function(){
		$("label[name='checkbox0']").each(function(){
				$(this).html($("#member0").children('option:selected').html().substr(5));
				});
		$("label[name='checkbox1']").each(function(){
				$(this).html($("#member1").children('option:selected').html().substr(5));
				});
		$("label[name='checkbox2']").each(function(){
				$(this).html($("#member2").children('option:selected').html().substr(5));
			});
	$("#member0").change(function(){
		$("label[name='checkbox0']").each(function(){
				$(this).html($("#member0").children('option:selected').html().substr(5));
			});
		});
	$("#member1").change(function(){
		$("label[name='checkbox1']").each(function(){
				$(this).html($("#member1").children('option:selected').html().substr(5));
			});
		});
	$("#member2").change(function(){
		$("label[name='checkbox2']").each(function(){
				$(this).html($("#member2").children('option:selected').html().substr(5));
			});
		});
	$("#checkbox_label").click(function(){
		$(this).prev("span").children("input").attr("checked", "true");
		});
	$("#radio_label").click(function(){
		$(this).prev("span").children("input").attr("checked", "true");
		});
});
</script>

<div class="container-fluid">
	<div class="row-fluid">
		<div class="span12">
			<form class="form-horizontal" id="table" method="post" action="#">
				<fieldset>
				    <legend>Edit Team Information</legend> 

					<div class="row">
						<div class="col-md-3">
							<div class="form-group">
								<div class="col-md-4">
                    				<label class="control-label">Team Name</label>
								</div>
								<div class="col-md-8">
                    				<input class="form-control" name="name" type="text" value="{{.form_pre_information.Name}}"/> 
								</div>
							</div>

							<div class="form-group">
								<div class="col-md-4">
                    				<label class="control-label">Team Rank</label>
								</div>
								<div class="col-md-8">
									<input class="form-control" name="rank" type="text" value="{{.form_pre_information.Rank}}"/> 
								</div>
							</div>

							{{range $key, $val := .form_pre_member}}
								<div class="form-group">
									<div class="col-md-4">
										<label class="control-label"> Member </label>
									</div>
									<div class="col-md-8">
										<select id="member{{$key}}" name="member{{$key}}" class="form-control selectpicker" value="{{$val}}">
											<option value="0" {{EqThenSelected $val 0 | Attr}}></option>
											{{range $in_key, $in_val := $.other_all_person}}
												<option value="{{$in_val.Id}}" {{EqThenSelected $val $in_val.Id | Attr}}>{{$in_val.Grade}}-{{$in_val.Name}}</option>
											{{end}}
										</select>
									</div>
								</div>
							{{end}}

							<div class="form-group">
								<div class="col-md-6">
                    				<button type="submit" class="form-control btn btn-large btn-success"><strong>Submit</strong></button>
								</div>
								<div class="col-md-6">
									<button type="button" class="form-control btn btn-large btn-warning" onclick="window.location='{{.url_current_table}}/edit'"><strong>Cancel</strong></button>
								</div>
							</div>
						</div>

						<div class="col-md-9">
							<table class="table">
								<tbody>
									{{ range $key, $val := .form_problem_status}}
										<tr>
											<td>
												<label class="form-control"> {{$val.Name}}  </label>
											</td>
											<td>
												<div class="input-group">
											    	<span class="input-group-addon">
														<input type="checkbox" value="1" name="problem_participant{{$key}}[]" {{EqThenChecked $val.Member1 1 | Attr}}>
													</span>
													<label class="form-control" id="checkbox_label" name="checkbox0" >Member1</label>
											    	<span class="input-group-addon">
														<input type="checkbox" value="2" name="problem_participant{{$key}}[]" {{EqThenChecked $val.Member2 1 | Attr}}>
													</span>
													<label class="form-control" id="checkbox_label" name="checkbox1" >Member2</label>
											    	<span class="input-group-addon">
														<input type="checkbox" value="4" name="problem_participant{{$key}}[]" {{EqThenChecked $val.Member3 1 | Attr}}>
													</span>
													<label class="form-control" id="checkbox_label" name="checkbox2" >Member3</label>
												</div>
											</td>
											<td>
												<div class="input-group">
											    	<span class="input-group-addon">
														<input type="radio"  value="0" name="problem_status{{$key}}" {{EqThenChecked $val.Status 0 | Attr}}>
													</span>
													<label class="form-control" id="radio_label">未ＡＣ</label>
													<span class="input-group-addon">
														<input type="radio"  value="1" name="problem_status{{$key}}" {{EqThenChecked $val.Status 1 | Attr}}>
													</span>
													<label class="form-control" id="radio_label">ＡＣ</label>
													<span class="input-group-addon">
														<input type="radio"  value="2" name="problem_status{{$key}}" {{EqThenChecked $val.Status 2 | Attr}}>
													</span>
													<label class="form-control" id="radio_label">赛后ＡＣ</label>
												</div>
											</td>
										</td>
									{{end}}
								</tbody>
							</table>
						</div>
					</div>
				</fieldset>
		    </form>
		</div>
	</div>
</div>
{{end}}
