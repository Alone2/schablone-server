from typing import Any, Dict

import httpx

from ...client import AuthenticatedClient
from ...types import UNSET, Response


def _get_kwargs(
    *,
    client: AuthenticatedClient,
    username: str,
    firstname: str,
    lastname: str,
    password: str,
) -> Dict[str, Any]:
    url = "{}/user/create".format(client.base_url)

    headers: Dict[str, str] = client.get_headers()
    cookies: Dict[str, Any] = client.get_cookies()

    params: Dict[str, Any] = {}
    params["username"] = username

    params["firstname"] = firstname

    params["lastname"] = lastname

    params["password"] = password

    params = {k: v for k, v in params.items() if v is not UNSET and v is not None}

    return {
        "method": "post",
        "url": url,
        "headers": headers,
        "cookies": cookies,
        "timeout": client.get_timeout(),
        "params": params,
    }


def _build_response(*, response: httpx.Response) -> Response[Any]:
    return Response(
        status_code=response.status_code,
        content=response.content,
        headers=response.headers,
        parsed=None,
    )


def sync_detailed(
    *,
    client: AuthenticatedClient,
    username: str,
    firstname: str,
    lastname: str,
    password: str,
) -> Response[Any]:
    """Create a new user

    Args:
        username (str):
        firstname (str):
        lastname (str):
        password (str):

    Returns:
        Response[Any]
    """

    kwargs = _get_kwargs(
        client=client,
        username=username,
        firstname=firstname,
        lastname=lastname,
        password=password,
    )

    response = httpx.request(
        verify=client.verify_ssl,
        **kwargs,
    )

    return _build_response(response=response)


async def asyncio_detailed(
    *,
    client: AuthenticatedClient,
    username: str,
    firstname: str,
    lastname: str,
    password: str,
) -> Response[Any]:
    """Create a new user

    Args:
        username (str):
        firstname (str):
        lastname (str):
        password (str):

    Returns:
        Response[Any]
    """

    kwargs = _get_kwargs(
        client=client,
        username=username,
        firstname=firstname,
        lastname=lastname,
        password=password,
    )

    async with httpx.AsyncClient(verify=client.verify_ssl) as _client:
        response = await _client.request(**kwargs)

    return _build_response(response=response)
