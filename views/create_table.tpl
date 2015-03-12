{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form id="table" method="post" action="/create_table">
				<fieldset>
				    <legend>Create Table</legend> 

                    <label>Table Title</label>
                    <input name="title" type="text" /> 
                    <span class="help-block">the title in hust</span> 

                    <label>Problem Number</label>
                    <input name="number" type="text" /> 
                    <span class="help-block">the title in hust</span> 

                    <label>Problem Source</label>
                    <input name="source" type="text" /> 
                    <span class="help-block">the title in hust</span> 

                    <button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
