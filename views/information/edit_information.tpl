{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Edit Information - SYSU Train System</title>
{{end}}

{{define "body"}}
<script type="text/javascript">
$(document).ready(function(){
		$("label[name='checkbox0']").each(function(){
				$(this).html($("#member0").children('option:selected').html());
				});
		$("label[name='checkbox1']").each(function(){
				$(this).html($("#member1").children('option:selected').html());
				});
		$("label[name='checkbox2']").each(function(){
				$(this).html($("#member2").children('option:selected').html());
			});
	$("#member0").change(function(){
		$("label[name='checkbox0']").each(function(){
				$(this).html($("#member0").children('option:selected').html());
			});
		});
	$("#member1").change(function(){
		$("label[name='checkbox1']").each(function(){
				$(this).html($("#member1").children('option:selected').html());
			});
		});
	$("#member2").change(function(){
		$("label[name='checkbox2']").each(function(){
				$(this).html($("#member2").children('option:selected').html());
			});
		});
	$("#checkbox_label").click(function(){
		$(this).siblings("input").attr("checked", "true");
		});
	$("#radio_label").click(function(){
		$(this).siblings("input").attr("checked", "true");
		});
});
</script>

<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<form id="table" method="post" action="#">
				<fieldset>
				    <legend>Edit Team Information</legend> 

					<div class="row">
						<div class="col-md-4">
                    		<label>Team name</label>
                    		<input name="name" type="text" value="{{.form_pre_information.Name}}"/> 
                    		<span class="help-block">team name</span> 

                    		<label>Team Rank</label>
                    		<input name="rank" type="text" value="{{.form_pre_information.Rank}}"/> 
                    		<span class="help-block">team rank without other board</span> 

							{{range $key, $val := .form_pre_member}}
								<label> Member </label>
								<select id="member{{$key}}" name="member{{$key}}" class="selectpicker" value="{{$val}}">
									<option value="0" {{EqThenSelected $val 0 | Attr}}></option>
									{{range $in_key, $in_val := $.other_all_person}}
										<option value="{{$in_val.Id}}" {{EqThenSelected $val $in_val.Id | Attr}}>{{$in_val.Grade}}-{{$in_val.Name}}</option>
									{{end}}
								</select>
                    			<span class="help-block">team member</span> 
							{{end}}

                    		<button type="submit" class="btn btn-success">Submit</button>
							<button type="button" class="btn btn-warning" onclick="window.location='{{.url_current_table}}/edit'">Cancel</button>
						</div>

						<div class="col-md-8">
							<table class="table table-striped table-hover table-bordered">
								<tbody>
									{{ range $key, $val := .form_problem_status}}
										<tr>
											<td class="text-center">
												<label> {{$val.Name}} :  </label>
											</td>
											<td class="text-center">
												    <label class="checkbox-inline">
														<input type="checkbox" value="1" name="problem_participant{{$key}}[]" {{if compare $val.Member1 1}} checked="true" {{end}}><label id="checkbox_label" name="checkbox0" >我是１号</label>
													</label>
												    <label class="checkbox-inline">
														<input type="checkbox" value="2" name="problem_participant{{$key}}[]" {{if compare $val.Member2 1}} checked="true" {{end}}><label id="checkbox_label" name="checkbox1" >我是２号</label>
													</label>
												    <label class="checkbox-inline">
														<input type="checkbox" value="4" name="problem_participant{{$key}}[]" {{if compare $val.Member3 1}} checked="true" {{end}}><label id="checkbox_label" name="checkbox2" >我是３号</label>
													</label>
											</td>
											<td class="text-center">
													<label class="radio-inline">
														<input type="radio"  value="0" name="problem_status{{$key}}" {{if compare $val.Status 0}} checked="true" {{end}}><label id="radio_label">未ＡＣ</label>
													</label>
													<label class="radio-inline">
														<input type="radio"  value="1" name="problem_status{{$key}}" {{if compare $val.Status 1}} checked="true" {{end}}><label id="radio_label">ＡＣ</label>
													</label>
													<label class="radio-inline">
														<input type="radio"  value="2" name="problem_status{{$key}}" {{if compare $val.Status 2}} checked="true" {{end}}><label id="radio_label">赛后ＡＣ</label>
													</label>
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
