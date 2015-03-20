{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form id="table" method="post" action="/create_table">
				<fieldset>
				    <legend>Create Table</legend> 

                    			<label>Table Title</label>
                    			<input name="contest_name" type="text" value="{{.Post.ContestName}}"/> 
                    			<span class="help-block">the title in hust</span> 

                    			<label>Problem Number</label>
                    			<input name="problem_number" type="text" value="{{.Post.ProblemNumber}}"/> 
                    			<span class="help-block">the title in hust</span> 

                    			<label>Problem Source</label>
                    			<input name="source" type="text" value="{{.Post.Source}}"/> 
                    			<span class="help-block">the title in hust</span> 

		    			<input type="hidden" name="table_id" value="{{.Post.Id}}">
                    			<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
