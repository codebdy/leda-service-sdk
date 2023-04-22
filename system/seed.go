package system

const systemSeed = `
{
  "classes": [
    {
      "name": "Meta",
      "uuid": "META_ENTITY_UUID",
      "stereoType": "Entity",
      "root": true,
      "attributes": [
        {
          "name": "id",
          "primary": true,
          "type": "ID",
          "uuid": "META_COLUMN_ID_UUID"
        },
        {
          "name": "name",
          "type": "String",
          "unique": true,
          "nullable": true,
          "uuid": "META_COLUMN_NAME_UUID"
        },
        {
          "name": "content",
          "nullable": true,
          "type": "JSON",
          "uuid": "META_COLUMN_CONTENT_UUID"
        },
        {
          "name": "publishedContent",
          "nullable": true,
          "type": "JSON",
          "uuid": "META_COLUMN_PUBLISHED_CONTENT_UUID"
        },
        {
          "name": "publishedAt",
          "type": "Date",
          "nullable": true,
          "uuid": "META_COLUMN_PUBLISHED_AT_UUID"
        },
        {
          "name": "createdAt",
          "createDate": true,
          "type": "Date",
          "uuid": "META_COLUMN_CREATED_AT_UUID"
        },
        {
          "name": "updatedAt",
          "type": "Date",
          "updateDate": true,
          "uuid": "META_COLUMN_UPDATED_AT_UUID"
        }
      ]
    },
    {
      "name": "App",
      "uuid": "APP_ENTITY_UUID",
      "stereoType": "Entity",
      "root": true,
      "attributes": [
        {
          "name": "id",
          "primary": true,
          "type": "ID",
          "uuid": "APP_COLUMN_ID_UUID"
        },
        {
          "name": "name",
          "type": "String",
          "unique": true,
          "uuid": "APP_COLUMN_NAME_UUID"
        },
        {
          "name": "title",
          "nullable": true,
          "type": "String",
          "uuid": "APP_COLUMN_TITLE_UUID"
        },
        {
          "name": "imageUrl",
          "nullable": true,
          "type": "String",
          "uuid": "APP_COLUMN_IMAGE_URL_UUID"
        },
        {
          "name": "metaId",
          "nullable": true,
          "type": "ID",
          "uuid": "APP_COLUMN_META_ID_UUID"
        },
        {
          "name": "createdAt",
          "createDate": true,
          "type": "Date",
          "uuid": "APP_COLUMN_CREATED_AT_UUID"
        },
        {
          "name": "updatedAt",
          "type": "Date",
          "updateDate": true,
          "uuid": "APP_COLUMN_UPDATED_AT_UUID"
        }
      ]
    },
    {
      "name": "Service",
      "uuid": "SERVICE_ENTITY_UUID",
      "stereoType": "Entity",
      "root": true,
      "attributes": [
        {
          "name": "id",
          "primary": true,
          "type": "ID",
          "uuid": "SERVICE_COLUMN_ID_UUID"
        },
        {
          "name": "name",
          "type": "String",
          "unique": true,
          "uuid": "SERVICE_COLUMN_NAME_UUID"
        },
        {
          "name": "title",
          "nullable": true,
          "type": "String",
          "uuid": "SERVICE_COLUMN_TITLE_UUID"
        },
        {
          "name": "imageUrl",
          "nullable": true,
          "type": "String",
          "uuid": "SERVICE_COLUMN_IMAGE_URL_UUID"
        },
        {
          "name": "metaId",
          "nullable": true,
          "type": "ID",
          "uuid": "SERVICE_COLUMN_META_ID_UUID"
        },
        {
          "name": "isSystem",
          "type": "Boolean",
          "nullable": true,
          "uuid": "SERVICE_COLUMN_SYSTEM_UUID"
        },
        {
          "name": "createdAt",
          "createDate": true,
          "type": "Date",
          "uuid": "SERVICE_COLUMN_CREATED_AT_UUID"
        },
        {
          "name": "updatedAt",
          "type": "Date",
          "updateDate": true,
          "uuid": "SERVICE_COLUMN_UPDATED_AT_UUID"
        }
      ]
    }
  ]

}
`
