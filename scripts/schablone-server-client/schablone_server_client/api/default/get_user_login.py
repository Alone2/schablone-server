from typing import Any, Dict, Optional, Union, cast

import httpx

from ...client import AuthenticatedClient
from ...types import UNSET, Response


def _get_kwargs(
    *,
    client: AuthenticatedClient,
    username: str,
    password: str,
) -> Dict[str, Any]:
    url = "{}/user/login".format(client.base_url)

    headers: Dict[str, str] = client.get_headers()
    cookies: Dict[str, Any] = client.get_cookies()

    params: Dict[str, Any] = {}
    params["username"] = username

    params["password"] = password

    params = {k: v for k, v in params.items() if v is not UNSET and v is not None}

    return {
        "method": "get",
        "url": url,
        "headers": headers,
        "cookies": cookies,
        "timeout": client.get_timeout(),
        "params": params,
    }


def _parse_response(*, response: httpx.Response) -> Optional[Union[Any, str]]:
    if response.status_code == 200:
        response_200 = cast(str, response.json())
        return response_200
    if response.status_code == 400:
        response_400 = cast(Any, None)
        return response_400
    if response.status_code == 404:
        response_404 = cast(Any, None)
        return response_404
    return None


def _build_response(*, response: httpx.Response) -> Response[Union[Any, str]]:
    return Response(
        status_code=response.status_code,
        content=response.content,
        headers=response.headers,
        parsed=_parse_response(response=response),
    )


def sync_detailed(
    *,
    client: AuthenticatedClient,
    username: str,
    password: str,
) -> Response[Union[Any, str]]:
    """Get API key

    Args:
        username (str):
        password (str):

    Returns:
        Response[Union[Any, str]]
    """

    kwargs = _get_kwargs(
        client=client,
        username=username,
        password=password,
    )

    response = httpx.request(
        verify=client.verify_ssl,
        **kwargs,
    )

    return _build_response(response=response)


def sync(
    *,
    client: AuthenticatedClient,
    username: str,
    password: str,
) -> Optional[Union[Any, str]]:
    """Get API key

    Args:
        username (str):
        password (str):

    Returns:
        Response[Union[Any, str]]
    """

    return sync_detailed(
        client=client,
        username=username,
        password=password,
    ).parsed


async def asyncio_detailed(
    *,
    client: AuthenticatedClient,
    username: str,
    password: str,
) -> Response[Union[Any, str]]:
    """Get API key

    Args:
        username (str):
        password (str):

    Returns:
        Response[Union[Any, str]]
    """

    kwargs = _get_kwargs(
        client=client,
        username=username,
        password=password,
    )

    async with httpx.AsyncClient(verify=client.verify_ssl) as _client:
        response = await _client.request(**kwargs)

    return _build_response(response=response)


async def asyncio(
    *,
    client: AuthenticatedClient,
    username: str,
    password: str,
) -> Optional[Union[Any, str]]:
    """Get API key

    Args:
        username (str):
        password (str):

    Returns:
        Response[Union[Any, str]]
    """

    return (
        await asyncio_detailed(
            client=client,
            username=username,
            password=password,
        )
    ).parsed
