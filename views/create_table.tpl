{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form>
				<fieldset>
				    <legend>Create Table</legend> 

                    <label>Table Title</label>
                    <input type="text" /> 
                    <span class="help-block">the title in hust</span> 

                    <label>Problem Number</label>
                    <input type="text" /> 
                    <span class="help-block">the title in hust</span> 

                    <label>Problem Source</label>
                    <input type="text" /> 
                    <span class="help-block">the title in hust</span> 

                    <button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
