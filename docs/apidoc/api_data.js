define({ "api": [
  {
    "type": "post",
    "url": "/api/v1/course",
    "title": "add course",
    "version": "1.0.0",
    "name": "addCourse",
    "group": "Course",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>course name</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "course_id",
            "description": "<p>course id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "unit",
            "description": "<p>course unit</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "course",
            "description": "<p>course model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/cource.go",
    "groupTitle": "Course"
  },
  {
    "type": "delete",
    "url": "/api/v1/course/:id",
    "title": "delete course",
    "version": "1.0.0",
    "name": "deleteCourse",
    "group": "Course",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/cource.go",
    "groupTitle": "Course"
  },
  {
    "type": "put",
    "url": "/api/v1/course/:id",
    "title": "edit course",
    "version": "1.0.0",
    "name": "editCourse",
    "group": "Course",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>course name</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "course_id",
            "description": "<p>course id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "unit",
            "description": "<p>course unit</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "course",
            "description": "<p>course model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/cource.go",
    "groupTitle": "Course"
  },
  {
    "type": "get",
    "url": "/api/v1/course/:id",
    "title": "get one course",
    "version": "1.0.0",
    "name": "getCourse",
    "group": "Course",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "course",
            "description": "<p>course model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/cource.go",
    "groupTitle": "Course"
  },
  {
    "type": "get",
    "url": "/api/v1/course",
    "title": "get courses list",
    "version": "1.0.0",
    "name": "getCoursesList",
    "group": "Course",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>list page</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "limit",
            "description": "<p>list limit</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "course",
            "description": "<p>course model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/cource.go",
    "groupTitle": "Course"
  },
  {
    "type": "get",
    "url": "/api/v1/course/mycourse",
    "title": "Get my courses",
    "version": "1.0.0",
    "name": "myCource",
    "group": "Student",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "course",
            "description": "<p>course model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/cource.go",
    "groupTitle": "Student"
  },
  {
    "type": "post",
    "url": "/api/v1/student/login",
    "title": "Student login",
    "version": "1.0.0",
    "name": "studentLogin",
    "group": "Student",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "student_id",
            "description": "<p>unique student id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "student",
            "description": "<p>student model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/student.go",
    "groupTitle": "Student"
  },
  {
    "type": "post",
    "url": "/api/v1/student/register",
    "title": "Student register",
    "version": "1.0.0",
    "name": "studentRegister",
    "group": "Student",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "first_name",
            "description": "<p>first name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "last_name",
            "description": "<p>last name</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "student_id",
            "description": "<p>unique student id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "student",
            "description": "<p>student model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/student.go",
    "groupTitle": "Student"
  },
  {
    "type": "post",
    "url": "/api/v1/teacher/login",
    "title": "Teacher login",
    "version": "1.0.0",
    "name": "teacherLogin",
    "group": "Teacher",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "teacher_id",
            "description": "<p>unique teacher id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "teacher",
            "description": "<p>teacher model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/teacher.go",
    "groupTitle": "Teacher"
  },
  {
    "type": "post",
    "url": "/api/v1/teacher/register",
    "title": "Teacher register",
    "version": "1.0.0",
    "name": "teacherRegister",
    "group": "Teacher",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "first_name",
            "description": "<p>first name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "last_name",
            "description": "<p>last name</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "teacher_id",
            "description": "<p>unique teacher id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "teacher",
            "description": "<p>teacher model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/teacher.go",
    "groupTitle": "Teacher"
  },
  {
    "type": "post",
    "url": "/api/v1/presentation",
    "title": "add presentation",
    "version": "1.0.0",
    "name": "addPresentation",
    "group": "presentation",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "cid",
            "description": "<p>course id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "time",
            "description": "<p>example 12:30</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "capacity",
            "description": "<p>presentation capacity</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/presentation.go",
    "groupTitle": "presentation"
  },
  {
    "type": "post",
    "url": "/api/v1/presentation/:id",
    "title": "add to obtained",
    "version": "1.0.0",
    "name": "addToObtained",
    "group": "presentation",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>error message</p>"
          }
        ]
      }
    },
    "filename": "web/v1/obtained.go",
    "groupTitle": "presentation"
  }
] });
