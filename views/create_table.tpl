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
						<select name="problem_number" class="selectpicker" value="0">
							<option value="0" selected="selected"></option>
							<option value="1"> A-1</option>
							<option value="2"> B-2</option>
							<option value="3"> C-3</option>
							<option value="4"> D-4</option>
							<option value="5"> E-5</option>
							<option value="6"> F-6</option>
							<option value="7"> G-7</option>
							<option value="8"> H-8</option>
							<option value="9"> I-9</option>
							<option value="10">J-10</option>
							<option value="11">K-11</option>
							<option value="12">L-12</option>
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
