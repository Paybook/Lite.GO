//View widget.ejs
app.controller('accountsController',function($scope){

  $scope.getTransactions = function(id_account){
    window.location = window.location.protocol + "//" + window.location.host + "/transactions/view/" + id_account;
  };

});
