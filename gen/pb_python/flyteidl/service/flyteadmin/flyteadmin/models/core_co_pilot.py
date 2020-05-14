# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.co_pilot_metadata_format import CoPilotMetadataFormat  # noqa: F401,E501


class CoreCoPilot(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'input_path': 'str',
        'output_path': 'str',
        'format': 'CoPilotMetadataFormat'
    }

    attribute_map = {
        'input_path': 'input_path',
        'output_path': 'output_path',
        'format': 'format'
    }

    def __init__(self, input_path=None, output_path=None, format=None):  # noqa: E501
        """CoreCoPilot - a model defined in Swagger"""  # noqa: E501

        self._input_path = None
        self._output_path = None
        self._format = None
        self.discriminator = None

        if input_path is not None:
            self.input_path = input_path
        if output_path is not None:
            self.output_path = output_path
        if format is not None:
            self.format = format

    @property
    def input_path(self):
        """Gets the input_path of this CoreCoPilot.  # noqa: E501


        :return: The input_path of this CoreCoPilot.  # noqa: E501
        :rtype: str
        """
        return self._input_path

    @input_path.setter
    def input_path(self, input_path):
        """Sets the input_path of this CoreCoPilot.


        :param input_path: The input_path of this CoreCoPilot.  # noqa: E501
        :type: str
        """

        self._input_path = input_path

    @property
    def output_path(self):
        """Gets the output_path of this CoreCoPilot.  # noqa: E501


        :return: The output_path of this CoreCoPilot.  # noqa: E501
        :rtype: str
        """
        return self._output_path

    @output_path.setter
    def output_path(self, output_path):
        """Sets the output_path of this CoreCoPilot.


        :param output_path: The output_path of this CoreCoPilot.  # noqa: E501
        :type: str
        """

        self._output_path = output_path

    @property
    def format(self):
        """Gets the format of this CoreCoPilot.  # noqa: E501


        :return: The format of this CoreCoPilot.  # noqa: E501
        :rtype: CoPilotMetadataFormat
        """
        return self._format

    @format.setter
    def format(self, format):
        """Sets the format of this CoreCoPilot.


        :param format: The format of this CoreCoPilot.  # noqa: E501
        :type: CoPilotMetadataFormat
        """

        self._format = format

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreCoPilot, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreCoPilot):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other