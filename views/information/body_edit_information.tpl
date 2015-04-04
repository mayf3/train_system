{{define "information_body"}}
<div class="container-fluid">
	<div class="row-fluid">
		<div class="span12">
			<form class="form-horizontal" id="table" method="post" action="#">
				<fieldset>
				    <legend>
						<strong>
							Edit Team Information
						</strong>	
					</legend> 

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
                    				<button type="submit" class="form-control btn btn-success">
										<strong>
											Submit
										</strong>
									</button>
								</div>
								<div class="col-md-6">
									<button type="button" class="form-control btn btn-warning" onclick="window.location='{{.url_current_table}}/edit'">
										<strong>
											Back
										</strong>
									</button>
								</div>
							</div>
						</div>

						<div class="col-md-9">
							<table class="table">
								<tbody>
									{{ range $key, $val := .form_problem_status}}
										<tr>
											<td class="text_center">
												<button type="button" class="btn btn-danger">
													{{$val.Name}} 
												</button>
											</td>
											<td class="text-center">
												<input hidden="hidden" type="checkbox" value="1" name="problem_participant{{$key}}[]" {{EqThenChecked $val.Member1 1 | Attr}}>
												<button name="checked_button0" type="button" class="btn btn-primary">Member1</button>
											</td>
											<td class="text-center">
												<input hidden="hidden" type="checkbox" value="2" name="problem_participant{{$key}}[]" {{EqThenChecked $val.Member2 1 | Attr}}>
												<button name="checked_button1" type="button" class="btn btn-primary">Member2</button>
											</td>
											<td class="text-center">
												<input hidden="hidden" type="checkbox" value="4" name="problem_participant{{$key}}[]" {{EqThenChecked $val.Member3 1 | Attr}}>
												<button name="checked_button2" type="button" class="btn btn-primary">Member3</button>
											</td>
											<td class="text-center">
												<input hidden="hidden" type="radio"  value="0" name="problem_status{{$key}}" {{EqThenChecked $val.Status 0 | Attr}}>
												<button name="radio_button" type="button" class="btn btn-primary">未AC</button>
											</td>
											<td class="text-center">
												<input hidden="hidden" type="radio"  value="1" name="problem_status{{$key}}" {{EqThenChecked $val.Status 1 | Attr}}>
												<button name="radio_button" type="button" class="btn btn-primary">AC</button>
											</td>
											<td class="text-center">
												<input hidden="hidden" type="radio"  value="2" name="problem_status{{$key}}" {{EqThenChecked $val.Status 2 | Attr}}>
												<button name="radio_button" type="button" class="btn btn-primary">赛后AC</button>
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
