<div class="container">
    <div class="row">
        <div class="col-sm-12">
          <div class="panel panel-info">
            <div class="panel-heading">
              <h3 class="panel-title">Add Account</h3>
            </div>
            <div class="panel-body">
              <script type="text/javascript" src="https://www.paybook.com/lib/js/widget/widget.js"></script>
              <script type="text/javascript">
                // Uncoment to use development environment
                pbWidget.setDev();
                pbWidget.setToken('<<< .Token >>>');
                pbWidget.chooseBank();
              </script>
              <div id="paybook-container"></div>
            </div>
          </div>
        </div>
    </div>
    <div class="row" ng-controller="accountsController">
        <div class="col-sm-12">
          <div class="panel panel-primary">
            <div class="panel-heading">
              <h3 class="panel-title">View Accounts</h3>
            </div>
            <div class="panel-body">
              <div class="list-group">
                <<< range $index, $element := .Accounts >>>
                  <button type="button" class="list-group-item" ng-click="getTransactions('<<< $element.IDAccount >>>')">
                  <img src="<<< $.Host >>><<< $element.SiteAvatar >>>" class="site-avatar">
                  <<< $element.SiteName >>>.
                  <<< $element.Name >>>
                <<< end >>>
              </div>
            </div>
          </div>
        </div>
    </div>
</div>
