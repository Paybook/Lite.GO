// Declare app level module which depends on views, and components
var app = angular.module('pbApp', [
  'ngRoute',
  'ngTable'
]);

app.factory('Session', ['$http', function($http) {
    return $http.get('/session').then(function(result) {
        return result.data.status;
    });
}]);

app.directive('menu', ['Session', function(Session) {
  function menu(scope, element, attrs) {
    Session.then(function(response){
      if (response){
        element.append('<nav class="navbar navbar-default">' +
                      ' <div class="container-fluid">' +
                      '  <ul class="nav navbar-nav">' +
                      '   <li><a href="/dashboard">Dashboard</a></li>'+
                      '   <li><a href="/logout">Logout</a></li>' +
                      '  </ul>' +
                      ' </div>' +
                      '</nav>');
      }
    });
  }
  return {
    link: menu
  };
}]);

app.factory('dataTable',['$filter','NgTableParams','$http', function($filter,NgTableParams,$http){
  return {
    tableParams: function(controller, reference, id){
      return new NgTableParams({
          page : 1,
          count: 10
        }, {
          total: 0,
          getData: function($defer, params) {
            var request = {};
            // Limit
            request.limit = params.count();
            var query = "?limit=" + params.count();
            if (params.page() > 1){
              var skip = params.count() * (params.page() - 1);
              request.skipe = skip;
              query = query + "&skip=" + skip;
            }

            // Sort
            var sort = params.sorting();
            var keys = Object.keys(sort);
            if (keys.length > 0){
              var order = sort[keys[0]];
              sortSplit = keys[0].split(".");
              if (sortSplit.length > 0){
                keys[0] = sortSplit[0];
              }
              request.sort = keys[0] + " " + order.toUpperCase();
              query = query + "&sort=" + keys[0] + " " + order.toUpperCase();
            }

            // Filter
            var filter = params.filter();
            keys = Object.keys(filter);
            var where = "";
            if (keys[0] !== undefined){
              var field = keys[0];
              where = '{"' + field  + '":{"contains":"' + filter[keys[0]]  + '"}}';
              query = query + "&where=" + where;
              request.where = where;
            }

            if (id !== null){
              query = query + "&" + id;
              request.id = id;
            }

            $http.get("/" + controller + "/count/" + id).then(function(result,status){
              reference.total = result.data;
              params.total(reference.total);
            });

            var config = {
             params: request,
             headers : {'Accept' : 'application/json'}
            };

            $http.get("/" + controller, config).then(function(result,status){
              var orderedData = params.sorting() ? $filter('orderBy')(result.data, params.orderBy()) : result.data;
              total = params.total();
              $defer.resolve(result.data);
            });

            // $http({
            //   method: 'GET',
            //   url: "/" + controller,
            //   params: request
            // }).then(function(result,status){
            //   var orderedData = params.sorting() ? $filter('orderBy')(result.data, params.orderBy()) : result.data;
            //   total = params.total();
            //   $defer.resolve(result.data);
            // });

          }
        });
    },

    getTotal: function(){

      return total;
    }
  };
}]);
