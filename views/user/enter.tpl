{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Enter - SYSU Train System</title>
{{end}}

{{define "body" }}
        <div class="container">
           <form class="form-horizontal col-sm-6 col-sm-offset-2" method="post" action="#">
                <div id="legend" class="">
                    <legend class="">Enter</legend>
                </div>

                <div class="form-group">
                <!-- Text input-->
                    <label class="control-label col-sm-4" for="input01">Username : </label>
                    <div class="col-sm-8">
                        <input name="user[username]" placeholder="username" class="form-control" type="text">
                    </div>
                </div>

                <div class="form-group">
                <!-- Text input-->
                    <label class="control-label col-sm-4" for="input01">Password : </label>
                    <div class="col-sm-8">
                        <input name="user[password]" placeholder="password" class="form-control" type="password">
                    </div>
                </div>

				<div class="form-group">
                <!-- Text input-->
                    <label class="control-label col-sm-4"></label>
                    <div class="col-sm-8">
						<input type="checkbox" value="true" name="option[remember_login]" >
						<label> Remember me for 7 days</label>
                    </div>
                </div>
                
                <div class="form-group">
                    <!-- Button -->
                    <label class="control-label col-sm-4"></label>
                    <div class="col-sm-8">
                        <button type="submit" class="btn btn-success">Sign In</button>
                    </div>
                </div>
            </form>
        </div>
{{end}}
