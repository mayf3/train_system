{{template "head.tpl" .}}

<script type="text/javascript">
$(document).ready(function(){
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
	$("#member3").change(function(){
		$("label[name='checkbox3']").each(function(){
				$(this).html($("#member3").children('option:selected').html());
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
				    <legend>Edit Team</legend> 

                    <label>Team name</label>
                    <input name="name" type="text" value="{{.Init.Name}}"/> 
                    <span class="help-block">team name</span> 

                    <label>Team Rank</label>
                    <input name="rank" type="text" value="{{.Init.Rank}}"/> 
                    <span class="help-block">team rank without other board</span> 

                    <label>Member1</label>
					<select id="member1" name="member1" class="selectpicker" data-style="btn-primary" value="{{.Member1}}">
						<option value="0" selected="selected"></option>
						{{range $key, $val := .all_person}}
							<option value="{{$val.Id}}">{{$val.Name}}</option>
						{{end}}
					</select>
                    <span class="help-block">team member 1</span> 

                    <label>Member2</label>
					<select id="member2" name="member2" class="selectpicker" data-style="btn-primary" value="{{.Member2}}">
						<option value="0" selected="selected"></option>
						{{range $key, $val := .all_person}}
							<option value="{{$val.Id}}">{{$val.Name}}</option>
						{{end}}
					</select>
                    <span class="help-block">team member 2</span> 

                    <label>Member3</label>
					<select id="member3" name="member3" class="selectpicker" data-style="btn-primary" value="{{.Member3}}">
						<option value="0" selected="selected"></option>
						{{range $key, $val := .all_person}}
							<option value="{{$val.Id}}">{{$val.Name}}</option>
						{{end}}
					</select>
                    <span class="help-block">team member 3</span> 

					{{ range $key, $val := .problem_list}}
						<label> Problem {{$val.Name}} :  </label>
						<div class="form-group">
						    <label class="checkbox-inline">
								<input type="checkbox" name="problem_participant{{$key}}[]"><label id="checkbox_label" name="checkbox1" >我是１号</label>
							</label>
						    <label class="checkbox-inline">
								<input type="checkbox" name="problem_participant{{$key}}[]"><label id="checkbox_label" name="checkbox2" >我是２号</label>
							</label>
						    <label class="checkbox-inline">
								<input type="checkbox" name="problem_participant{{$key}}[]"><label id="checkbox_label" name="checkbox3" >我是３号</label>
							</label>
						</div>
						<div class="form-group">
							<label class="radio-inline">
								<input type="radio"  value="0" name="problem_status{{$key}}"><label id="radio_label">未ＡＣ</label>
							</label>
							<label class="radio-inline">
								<input type="radio"  value="1" name="problem_status{{$key}}"><label id="radio_label">ＡＣ</label>
							</label>
							<label class="radio-inline">
								<input type="radio"  value="2" name="problem_status{{$key}}"><label id="radio_label">赛后ＡＣ</label>
							</label>
						</div>
					{{end}}

                    <button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
		    </form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
