# coding: utf-8

"""
    FirstQuadrant API

    The FirstQuadrant API is used to interact with FirstQuadrant programmatically. We also have SDKs available (coming soon).  # noqa: E501

    The version of the OpenAPI document: 0.0.0
    Contact: inquiry@firstquadrant.ai
    Generated by: https://openapi-generator.tech
"""

from setuptools import setup, find_packages  # noqa: H301

NAME = "openapi-client"
VERSION = "1.0.0"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = [
    "certifi >= 14.5.14",
    "frozendict ~= 2.3.4",
    "python-dateutil ~= 2.7.0",
    "setuptools >= 21.0.0",
    "typing_extensions ~= 4.3.0",
    "urllib3 ~= 1.26.7",
]

setup(
    name=NAME,
    version=VERSION,
    description="FirstQuadrant API",
    author="OpenAPI Generator community",
    author_email="inquiry@firstquadrant.ai",
    url="",
    keywords=["OpenAPI", "OpenAPI-Generator", "FirstQuadrant API"],
    python_requires=">=3.7",
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    long_description="""\
    The FirstQuadrant API is used to interact with FirstQuadrant programmatically. We also have SDKs available (coming soon).  # noqa: E501
    """
)