console.log('Loading function');

var AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB.DocumentClient();
const tableName = process.env.TABLE_NAME || 'poc-items-node'

/**
 * Provide an event that contains the following keys:
 *
 *   - operation: one of the operations in the switch statement below
 *   - tableName: required for operations that interact with DynamoDB
 *   - payload: a parameter to pass to the operation being performed
 */
exports.handler = async (event, context) => {
    //console.log('Received event:', JSON.stringify(event, null, 2));

    var body = JSON.parse(event.body);

    var operation = body.operation;

    switch (operation) {
        case 'create':
            body.payload.TableName = tableName;
            await dynamo.put(body.payload).promise();
            return {'message': 'ok'}
            break;
        case 'echo':
            return body.payload;
            break;
        default:
            return({'message': `Unknown operation: ${operation}`});
    }
};