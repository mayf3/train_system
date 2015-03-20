{{template "head.tpl" .}}

<div class="container">
	<div class="row-fluid">
		<div class="span12">
			<form>
				<fieldset>
				    <legend>Edit Imformation</legend> 

                    <label>Team Rank</label>
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
