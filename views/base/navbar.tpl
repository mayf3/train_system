        <nav class="navbar navbar-inverse navbar-fixed-top">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="/">SYSU Train System</a>
                </div>
                <div class="navbar-collapse collapse">
                    <ul class="nav navbar-nav">
                        <li id="nav_index" class="active">
							<a href="/">Home</a>
						</li>
                    </ul>
                    <ul class="nav navbar-nav">
                        <li id="nav_person">
							<a href="/person">Person</a>
						</li>
                    </ul>
					{{if .username}}
						<ul class="nav navbar-nav navbar-right">
                    	    <li id="nav_Enter">
								<a href="/user/logout">Logout</a>
							</li>
                    	</ul>
						<ul class="nav navbar-nav navbar-right">
                        	<li id="nav_index">
								<a href="#">{{.username}}</a>
							</li>
						</ul>
					{{else}}
						<ul class="nav navbar-nav navbar-right">
                    	    <li id="nav_Register">
								<a href="/user/register">Register</a>
							</li>
                    	</ul>
						<ul class="nav navbar-nav navbar-right">
                    	    <li id="nav_Enter">
								<a href="/user/enter">Enter</a>
							</li>
                    	</ul>
					{{end}}
                </div>
            </div>
        </nav>
