from __future__ import print_function

import boto3
import json
import os


print('Loading function')


def handler(event, context):
    '''Provide an event that contains the following keys:

      - operation: one of the operations in the operations dict below
      - payload: a parameter to pass to the operation being performed
    '''
    print("Received event: " + json.dumps(event))
    
    body = json.loads(event['body'])

    operation = body['operation']
    payload = body['payload']
   
    if operation == 'create':
        dynamodb = boto3.client('dynamodb')
        item = payload['Item']
        newItem = { 'id': {}, 'year': {} }

        newItem['id']['S'] = item['id']
        newItem['year']['S'] = item['year']
        print("Received item: " + json.dumps(item))

        tableName = os.getenv('TABLE_NAME', 'poc-items-python')
        dynamodb.put_item(TableName=tableName, Item=newItem)
        response = {
            'message': 'ok'
        }
        return response
    if operation == 'echo':
        return payload
    
    raise ValueError('Unrecognized operation "{}"'.format(operation))