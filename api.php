/**
 * @api {post} /users/ Create new user
 * @apiGroup Users
 * @apiParam {String} email User email
 * @apiParam {String} password User password
 *
 * @apiSuccess {Number} id user id
 * @apiSuccessExample Success-Response:
 *    HTTP/1.1 200 OK
 *    {
 *      "id": 11,
 *    }
 * @apiError EmailExistOrNotValid The id of the User not found
 *
 * @apiErrorExample Error-Response:
 *    HTTP/1.1 404 Data Not Valid
 *    {
 *      "error": "DataNotValid"
 *    }
 */

/**
 * @api {put} /users/ Update new user
 * @apiGroup Users
 * @apiParam {String} name User name
 * @apiParam {String} email User email
 * @apiParam {String} password User password
 *
 * @apiSuccess {Number} id user id
 * @apiSuccessExample Success-Response:
 *    HTTP/1.1 200 OK
 *    {
 *      "id": 11,
 *      "name": "Ivanov",
 *      "email": "ivanov@gmail.com"
 *    }
 * @apiError EmailExistOrNotValid User not update
 *
 * @apiErrorExample Error-Response:
 *    HTTP/1.1 404 User Not Update
 *    {
 *      "error": "UserNotUpdate"
 *    }
 */

/**
 * @api {delete} /users/:id Delete a User
 * @apiGroup Users
 * @apiParam {id} id product id

 * @apiSuccess {id} id user id
 * @apiSuccessExample Success-Response:
 *    HTTP/1.1 200 OK Remove User
 *    {
 *      "id": 11,
 *    }
 * @apiError EmailExistOrNotValid The User not delete
 *
 * @apiErrorExample Error-Response:
 *    HTTP/1.1 404 User Not Delete
 *    {
 *      "error": "UserNotFound"
 *    }
 */

/**
 * @api {post} /product/ Create a product
 * @apiGroup product
 * @apiPermission Administrator
 * @apiParam {String} name Product name
 * @apiParam {String} description Product description
 * @apiParam {Number} quantity Product quantity
 * @apiParam {String} composition Product composition
 * @apiParam {Number} price Product price

 * @apiSuccess {Number} id user id
 * @apiSuccess {String} name Product name
 * @apiSuccess {String} description Product description
 * @apiSuccess {Number} quantity Product quantity
 * @apiSuccess {String} composition Product composition
 * @apiSuccess {Number} price Product price

 * @apiSuccessExample Success-Response:
 *    HTTP/1.1 200 OK
 *    {
 *      "id": "1111"
 *      "name": "Pen",
 *      "description": "for drawing and drawings",
 *      "quantity": 1001,
 *      "composition": "graphite lead and wood",
 *      "price": 52
 *    }
 * @apiError DataNotValid The data of the product not valid
 *
 * @apiErrorExample Error-Response:
 *    HTTP/1.1 404 Data Not Valid
 *    {
 *      "error": "DataNotValid"
 *    }
 */

/**
 * @api {get} /product/:id Find a product
 * @apiGroup product
 * @apiParam {id} id product id
 * @apiSuccess {Number} id Product id
 * @apiSuccess {String} name Product name
 * @apiSuccess {String} description Product description
 * @apiSuccess {Number} quantity Product quantity
 * @apiSuccess {String} composition Product composition
 * @apiSuccess {Number} price Product price
 * @apiSuccessExample {json} Success
 *    HTTP/1.1 200 OK
 *    {
 *      "id": 10,
 *      "name": "Pencil",
 *      "description": "for drawing",
 *      "quantity": 1001, 
 *      "composition": "graphite lead and wood",
 *      "price": 54
 *    }
 * @apiErrorExample {json} Product not found
 *    HTTP/1.1 404 Product Not Found
 * @apiErrorExample {json} Find error
 *    HTTP/1.1 500 Internal Server Error
 */

/**
 * @api {get} /product Find all product
 * @apiGroup product
 * @apiSuccess {Number} id product id
 * @apiSuccess {String} name product name
 * @apiSuccessExample {json} Success
 *    HTTP/1.1 200 OK
 *    {
 *      "id": 1,
 *      "name": "Pen"
 *    }
 * @apiErrorExample {json} product not found
 *    HTTP/1.1 404 Product Not Found
 * @apiErrorExample {json} Find error
 *    HTTP/1.1 500 Internal Server Error
 */

/**
 * @api {put} /product/:id Update a product
 * @apiGroup product
 * @apiPermission Administrator
 * @apiParam {id} id product id
 * @apiParam {String} name Task name
 * @apiParam {String} description Product description
 * @apiParam {Number} quantity Product quantity
 * @apiParam {String} composition Product composition
 * @apiParam {Number} price Product price

 * @apiParamExample {json} Input
 *    {
 *      "name": "Pen",
 *      "description": "for drawing and drawings",
 *      "quantity": 1001,
 *      "composition": "graphite lead and wood",
 *      "price": 52
 *    }
 * @apiSuccessExample {json} Success
 *    HTTP/1.1 204 Success Update Product
 * @apiErrorExample {json} Update error
 *    HTTP/1.1 500 Internal Server Error
 */

/**
 * @api {delete} /product/:id Delete a product
 * @apiGroup product
 * @apiPermission Administrator
 * @apiParam {id} id product id
 * @apiSuccess {id} id product id
 * @apiSuccessExample Success-Response:
 *    HTTP/1.1 200 OK
 *    {
 *      "id": 11,
 *    }
 * @apiError EmailExistOrNotValid The id of the product not found
 *
 * @apiErrorExample Error-Response:
 *    HTTP/1.1 404 Product Not Found
 *    {
 *      "error": "ProductNotFound"
 *    }
 */
