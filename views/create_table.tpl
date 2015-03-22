{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form id="table" method="post" action="#">
				<fieldset>
				    <legend>Create Table</legend> 

                    	<label>Table Title</label>
                    	<input name="contest_name" type="text" value="{{.Init.ContestName}}"/> 
                    	<span class="help-block">SYSU-1, SYSU-2, SYU-3.....</span> 

                    	<label>Problem Number</label>
						<select name="problem_number" class="selectpicker">
							<option value="0" selected="selected"></option>
							<option value="1" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 1}} selected="selected" {{end}}  {{end}}> A-1</option>
							<option value="2" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 2}} selected="selected" {{end}}  {{end}}> B-2</option>
							<option value="3" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 3}} selected="selected" {{end}}  {{end}}> C-3</option>
							<option value="4" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 4}} selected="selected" {{end}}  {{end}}> D-4</option>
							<option value="5" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 5}} selected="selected" {{end}}  {{end}}> E-5</option>
							<option value="6" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 6}} selected="selected" {{end}}  {{end}}> F-6</option>
							<option value="7" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 7}} selected="selected" {{end}}  {{end}}> G-7</option>
							<option value="8" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 8}} selected="selected" {{end}}  {{end}}> H-8</option>
							<option value="9" {{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 9}} selected="selected" {{end}}  {{end}}> I-9</option>
							<option value="10"{{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 10}} selected="selected" {{end}} {{end}}>J-10</option>
							<option value="11"{{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 11}} selected="selected" {{end}} {{end}}>K-11</option>
							<option value="12"{{if .Init.ProblemNumber }} {{if eq .Init.ProblemNumber 12}} selected="selected" {{end}} {{end}}>L-12</option>
						</select>
                    	<span class="help-block">total problem number</span> 

                    	<label>Problem Source</label>
                    	<input name="source" type="text" value="{{.Init.Source}}"/> 
                    	<span class="help-block">for example : Regionals 2014 >> Asia - Dhaka</span> 

		    			<input type="hidden" name="table_id" value="{{.Init.Id}}">
						<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
