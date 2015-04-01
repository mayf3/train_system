{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Setting Table- SYSU Train System</title>
{{end}}

{{define "body"}}
<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form id="table" method="post" action="#">
				<fieldset>
				    <legend>Create Table</legend> 
                    	<label>Table Title</label>
                    	<input name="contest_name" type="text" value="{{.form_pre_table.ContestName}}"/> 
                    	<span class="help-block">SYSU-1, SYSU-2, SYU-3.....</span> 

                    	<label>Problem Number</label>
						<select name="problem_number" class="selectpicker" value="{{$.form_problem_number}}">
							{{range $key, $val := .base_problem_list}}
								<option value="{{$key}}" {{EqThenSelected $key $.form_problem_number | Attr}}> {{$val}}</option>
							{{end}}
						</select>
                    	<span class="help-block">total problem number</span> 

                    	<label>Problem Source</label>
                    	<input name="source" type="text" value="{{.form_pre_table.Source}}"/> 
                    	<span class="help-block">for example : Regionals 2014 >> Asia - Dhaka</span> 

		    			<input type="hidden" name="table_id" value="{{.form_pre_table.Id}}">
						<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>
{{end}}
