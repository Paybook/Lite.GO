<div ng-controller="transactionsController" ng-init="id_account='<<< .IDAccount >>>';">
  <img src="<<< .Host >>><<< .Account.SiteAvatar >>>" class="site-avatar">
  <<< .Account.SiteName >>>
  <<< .Account.Name >>>
  <h4 class="page-header">Transactions: </h4>
  <table ng-table="tableParams" class="table table-responsive table-inverse set-bg">
    <tr ng-repeat="row in $data">
      <td data-title="'Date'" sortable="'dt_transaction'">{{row.DTTransaction | date : format : timezone}}</td>
      <td data-title="'Description'" sortable="'description'" filter="{description: 'text'}">{{row.Description}}</td>
      <td data-title="'Amount'" sortable="'amount'">{{row.Amount | currency : symbol : fractionSize}}</td>
    </tr>
  </table>
</div>
